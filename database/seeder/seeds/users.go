package seeds

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/oullin/database"
	"github.com/oullin/pkg"
	"github.com/oullin/pkg/gorm"
	"strings"
	"time"
)

type UsersSeed struct {
	db *database.Connection
}

type UsersAttrs struct {
	Username string
	Name     string
	IsAdmin  bool
}

func MakeUsersSeed(db *database.Connection) *UsersSeed {
	return &UsersSeed{
		db: db,
	}
}

func (s UsersSeed) Create(attrs UsersAttrs) (database.User, error) {
	pass, _ := pkg.MakePassword("password")

	fake := database.User{
		UUID:         uuid.NewString(),
		FirstName:    attrs.Name,
		LastName:     "Tester",
		Username:     attrs.Username,
		DisplayName:  fmt.Sprintf("%s User", attrs.Name),
		Email:        fmt.Sprintf("%s@test.com", strings.Trim(attrs.Username, " ")),
		PasswordHash: pass.GetHash(),
		PublicToken:  uuid.NewString(),
		IsAdmin:      attrs.IsAdmin,
		Bio:          "Software Engineer with an eye for details.",
		VerifiedAt:   time.Now(),
	}

	result := s.db.Sql().Create(&fake)

	if gorm.HasDbIssues(result.Error) {
		return database.User{}, fmt.Errorf("issues creating users: %s", result.Error)
	}

	return fake, nil
}
