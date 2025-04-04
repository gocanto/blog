package main

import (
	baseValidator "github.com/go-playground/validator/v10"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/proxy"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log/slog"
	"net/http"
)

var environment *env.Environment
var validator *proxy.Validator

func init() {
	val := proxy.MakeValidatorFrom(baseValidator.New(
		baseValidator.WithRequiredStructEnabled(),
	))

	values, err := godotenv.Read("./../.env")

	if err != nil {
		panic("invalid .env file: " + err.Error())
	}

	environment = MakeEnv(values, val)
	validator = val
}

func main() {
	orm := MakeORM(environment)
	logs := MakeLogs(environment)
	adminUser := MakeAdminUser(environment)

	defer (*logs).Close()
	defer (*orm).Close()

	mux := http.NewServeMux()

	app := MakeApp(mux, &App{
		Validator: validator,
		Logs:      logs,
		Orm:       orm,
		AdminUser: adminUser,
		Env:       environment,
		Mux:       mux,
	})

	app.RegisterUsers()

	(*orm).Ping()
	slog.Info("Starting new server on :" + environment.Network.HttpPort)

	if err := http.ListenAndServe(environment.Network.GetHostURL(), mux); err != nil {
		slog.Error("Error starting server", "error", err)
		panic("Error starting server." + err.Error())
	}
}
