package media

import (
	"errors"
	"github.com/google/uuid"
	"path/filepath"
	"strings"
)

const Dir = "media"
const UsersDir = "users"
const PostsDir = "posts"
const StorageDir = "storage"

var maxFileSize = int64(50 * 1024 * 1024) // 50 MB in bytes

type Media struct {
	file   []byte
	header string
	ext    string
	name   string
	path   string
	size   int64
}

func MakeMedia(file []byte, headerName string) (*Media, error) {
	size := int64(len(file))
	ext := strings.ToLower(filepath.Ext(headerName))
	filename := uuid.New().String() + ext

	if len(file) <= 0 {
		return nil, errors.New("the given file is empty")
	}

	if size > maxFileSize {
		return nil, errors.New("the given file is too big")
	}

	if hasInvalidExt(ext) {
		return nil, errors.New("the given file type is invalid")
	}

	media := Media{
		file:   file,
		header: headerName,
		ext:    ext,
		name:   filename,
		size:   size,
	}

	media.path = filepath.Join(
		media.GetUsersImagesDir(),
		filename,
	)

	return &media, nil
}

func hasValidExt(ext string) bool {
	var allowedExtensions = []string{".jpg", ".jpeg", ".png"}

	for _, fileExt := range allowedExtensions {

		if fileExt == ext {
			return true
		}

	}

	return false
}

func hasInvalidExt(ext string) bool {
	return !hasValidExt(ext)
}
