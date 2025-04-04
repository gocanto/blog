package main

import (
	"github.com/gocanto/blog/app/database"
	"github.com/gocanto/blog/app/env"
	"github.com/gocanto/blog/app/logs"
	"github.com/gocanto/blog/app/proxy"
	"github.com/gocanto/blog/app/users"
	"strconv"
	"strings"
)

func MakeORM(env *env.Environment) *database.Orm {
	dbConn, err := database.MakeORM(env)

	if err != nil {
		panic("DB: error connecting to PostgreSQL: " + err.Error())
	}

	return dbConn
}

func MakeLogs(env *env.Environment) *logs.Driver {
	lDriver, err := logs.MakeFilesLogs(env)

	if err != nil {
		panic("Logs: error opening logs file: " + err.Error())
	}

	return &lDriver
}

func MakeAdminUser(env *env.Environment) *users.AdminUser {
	return &users.AdminUser{
		PublicToken:  env.App.AppUserAmin.PublicToken,
		PrivateToken: env.App.AppUserAmin.PrivateToken,
	}
}

func MakeEnv(values map[string]string, validate *proxy.Validator) *env.Environment {
	errorSufix := "Environment: "

	port, _ := strconv.Atoi(values["ENV_DB_PORT"])

	userAminEnvValues := &env.AppUserAminEnvValues{
		PublicToken:  strings.Trim(values["ENV_APP_ADMIN_PUBLIC_TOKEN"], " "),
		PrivateToken: strings.Trim(values["ENV_APP_ADMIN_PRIVATE_TOKEN"], " "),
	}

	app := env.AppEnvironment{
		Name:        values["ENV_APP_NAME"],
		Type:        values["ENV_APP_ENV_TYPE"],
		AppUserAmin: userAminEnvValues,
	}

	db := env.DBEnvironment{
		UserName:     values["ENV_DB_USER_NAME"],
		UserPassword: values["ENV_DB_USER_PASSWORD"],
		DatabaseName: values["ENV_DB_DATABASE_NAME"],
		Port:         port,
		Host:         values["ENV_DB_HOST"],
		DriverName:   "postgres",
		BinDir:       values["EN_DB_BIN_DIR"],
		URL:          values["ENV_DB_URL"],
		SSLMode:      values["ENV_DB_SSL_MODE"],
		TimeZone:     values["ENV_DB_TIMEZONE"],
	}

	logs := env.LogsEnvironment{
		Level:      values["ENV_APP_LOG_LEVEL"],
		Dir:        values["ENV_APP_LOGS_DIR"],
		DateFormat: values["ENV_APP_LOGS_DATE_FORMAT"],
	}

	net := env.NetEnvironment{
		HttpHost: values["ENV_HTTP_HOST"],
		HttpPort: values["ENV_HTTP_PORT"],
	}

	if _, err := validate.Rejects(app); err != nil {
		panic(errorSufix + "invalid [APP] model: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(db); err != nil {
		panic(errorSufix + "invalid [DB] model: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(userAminEnvValues); err != nil {
		panic(errorSufix + "invalid [AppUserAminEnvValues] model: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(logs); err != nil {
		panic(errorSufix + "invalid [LOGS] model: " + validate.GetErrorsAsJason())
	}

	if _, err := validate.Rejects(net); err != nil {
		panic(errorSufix + "invalid [NETWORK] model: " + validate.GetErrorsAsJason())
	}

	blog := &env.Environment{
		App:     app,
		DB:      db,
		Logs:    logs,
		Network: net,
	}

	if _, err := validate.Rejects(blog); err != nil {
		panic(errorSufix + "invalid blog [ENVIRONMENT] model: " + validate.GetErrorsAsJason())
	}

	return blog
}
