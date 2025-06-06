package middleware

import (
	"github.com/gocanto/blog/env"
	"github.com/gocanto/blog/pkg"
)

type MiddlewaresStack struct {
	env               *env.Environment
	middleware        []Middleware
	userAdminResolver func(seed string) bool
}

type Middleware func(pkg.BaseHandler) pkg.BaseHandler

func MakeMiddlewareStack(env *env.Environment, userAdminResolver func(seed string) bool) *MiddlewaresStack {
	return &MiddlewaresStack{
		env:               env,
		userAdminResolver: userAdminResolver,
		middleware:        []Middleware{},
	}
}
