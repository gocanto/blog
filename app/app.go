package main

import (
	"github.com/gocanto/blog/app/controller"
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/logs"
	"github.com/gocanto/blog/app/proxy"
	"github.com/gocanto/blog/app/users"
	"net/http"
)

type App struct {
	Validator *proxy.Validator `validate:"required"`
	Logs      *logs.Driver     `validate:"required"`
	Orm       *database.Orm    `validate:"required"`
	AdminUser *users.AdminUser `validate:"required"`
	Env       *env.Environment `validate:"required"`
	Mux       *http.ServeMux   `validate:"required"`
}

func MakeApp(mux *http.ServeMux, app *App) *App {
	app.Mux = mux

	return app
}

func (app App) RegisterUsers() {
	stack := controller.MakeMiddlewareStack(app.Env, func(seed string) bool {
		return app.AdminUser.IsAllowed(seed)
	})

	userController := users.UserController{
		Repository: users.MakeRepository(app.Orm, app.AdminUser),
		Validator:  app.Validator,
	}

	app.Mux.HandleFunc("POST /users", controller.CreateHandle(
		stack.Push(
			userController.Create,
			stack.Logging,
			stack.AdminUser,
		),
	))
}
