package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocanto/blog/app/controller"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/media"
	"io"
	"log/slog"
	"net/http"
)

type CreateRequestBag struct {
	FirstName            string `json:"first_name" validate:"required,min=4,max=250"`
	LastName             string `json:"last_name" validate:"required,min=4,max=250"`
	Username             string `json:"username" validate:"required,alphanum,min=4,max=50"`
	DisplayName          string `json:"display_name" validate:"omitempty,min=3,max=255"`
	Email                string `json:"email" validate:"required,email,max=250"`
	Password             string `json:"password" validate:"required,min=8"`
	PublicToken          string `json:"public_token"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required,eqfield=Password"`
	Bio                  string `json:"bio" validate:"omitempty"`
	ProfilePictureURL    string `json:"profile_picture_url" validate:"omitempty,url,max=2048"`
}

func (uc UserController) Create(w http.ResponseWriter, r *http.Request) *controller.HttpError {
	body, err := io.ReadAll(r.Body)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("Issue closing the request body", err)
		}
	}(r.Body)

	if err != nil {
		return controller.BadRequest("Invalid request payload: cannot read body", err)
	}

	fmt.Println("Raw body length:", len(body))
	// Reset the request body so it can be read again.
	r.Body = io.NopCloser(bytes.NewBuffer(body))

	// Get the multipart reader.
	mr, err := r.MultipartReader()
	if err != nil {
		return controller.BadRequest("Error getting multipart reader", err)
	}

	var fileBytes []byte
	var dataBytes []byte
	var fileHeaderName string

	for {
		part, err := mr.NextPart()

		if err == io.EOF {
			break
		}

		if err != nil {
			return controller.BadRequest("Error reading multipart parts", err)
		}

		// Check which part we got.
		switch part.FormName() {

		case "data":
			// Ensure this is a text field (not a file).
			if part.FileName() != "" {
				return controller.BadRequest("Expected 'data' to be a JSON text field", err)
			}

			dataBytes, err = io.ReadAll(part)
			if err != nil {
				return controller.BadRequest("Error reading data field", err)
			}

			fmt.Println("Received data field:", string(dataBytes))

		case "profile_picture_url":

			fileBytes, err = io.ReadAll(part)
			if err != nil {
				return controller.BadRequest("Error reading file", err)
			}
			fileHeaderName = part.FileName()
			fmt.Printf("Received file part: %d bytes\n", len(fileBytes))
		default:
			fmt.Println("Ignoring unexpected part:", part.FormName())
		}

		if err = part.Close(); err != nil {
			slog.Error("Issue closing the multi-part reader", err)
		}
	}

	// --- Save the file using fileBytes ---
	profilePic, err := media.MakeMedia(fileBytes, fileHeaderName)

	if err != nil {
		return controller.BadRequest("Error handling the given file", err)
	}

	if err := profilePic.Write(); err != nil {
		return controller.BadRequest("Error saving the given file", err)
	}

	var requestBag CreateRequestBag
	if err = json.Unmarshal(dataBytes, &requestBag); err != nil {
		return controller.BadRequest("Invalid request payload: malformed JSON", err)
	}

	validate := uc.Validator
	if rejects, err := validate.Rejects(requestBag); rejects {
		return controller.RespondWithErrors("Validation failed", validate.GetErrors(), err)
	}

	if result := uc.Repository.FindByUserName(requestBag.Username); result != nil {
		return controller.RespondWithErrors(
			fmt.Sprintf("user '%s' already exists", requestBag.Username),
			map[string]any{},
			nil,
		)
	}

	requestBag.PublicToken = r.Header.Get(env.ApiKeyHeader)
	created, err := uc.Repository.Create(requestBag)

	if err != nil {
		return controller.InternalServerError(err.Error(), err)
	}

	payload := map[string]any{
		"message": "User created successfully!",
		"user":    map[string]string{"uuid": created.UUID},
		//"data":    json.RawMessage(body),
	}

	return controller.SendJSON(w, http.StatusCreated, payload)
}
