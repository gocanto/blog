package controller

import (
	"github.com/gocanto/blog/app/env"
)

type MiddlewareStack struct {
	env               *env.Environment
	middleware        []Middleware
	userAdminResolver func(seed string) bool
}

type Middleware func(BaseController) BaseController

func MakeMiddlewareStack(env *env.Environment, userAdminResolver func(seed string) bool) *MiddlewareStack {
	return &MiddlewareStack{
		env:               env,
		userAdminResolver: userAdminResolver,
		middleware:        []Middleware{},
	}
}
