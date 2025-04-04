package media

import (
	"os"
	"path/filepath"
)

func (m *Media) Write() error {
	err := os.WriteFile(m.path, m.file, 0644)

	if err != nil {
		return err
	}

	return nil
}

func (m *Media) FileName() string {
	return m.name
}
func (m *Media) FilePath() string {
	return m.path
}

func (m *Media) HeaderName() string {
	return m.header
}

func (m *Media) Extension() string {
	return m.ext
}

func (m *Media) Filename() string {
	return m.name
}

func (m *Media) GetStorageDir() string {
	dir, err := os.Getwd()
	folder := StorageDir

	if err != nil {
		// Handle the error appropriately.
		// Path default to relative path if getting WD fails
		return "./" + folder
	}

	// Resolve the path at runtime.
	return filepath.Join(dir, StorageDir)
}

func (m *Media) GetMediaDir() string {
	return m.GetStorageDir() + "/" + Dir
}

func (m *Media) GetUsersImagesDir() string {
	return m.GetMediaDir() + "/" + UsersDir
}

func (m *Media) GetPostsImagesDir() string {
	return m.GetMediaDir() + "/" + PostsDir
}
