package user

import (
	"github.com/gocanto/blog/database"
	"github.com/gocanto/blog/packages/model"
)

type Repository struct {
	Connection database.Connection
}

func (r Repository) Create(user model.User) (error, model.User) {
	return nil, user
}
