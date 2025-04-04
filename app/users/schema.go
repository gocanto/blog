package users

import (
	"github.com/gocanto/blog/app/proxy"
)

type UserController struct {
	Validator  *proxy.Validator
	Repository *Repository
}

type CreatedUser struct {
	UUID string `json:"uuid"`
}
