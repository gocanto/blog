package env

const local = "local"
const staging = "staging"
const production = "production"
const ApiKeyHeader = "X-API-Key"

type AppEnvironment struct {
	Name        string                `validate:"required,min=4"`
	Type        string                `validate:"required,lowercase,oneof=local production staging"`
	AppUserAmin *AppUserAminEnvValues `validate:"required"`
}

type AppUserAminEnvValues struct {
	PublicToken  string `validate:"required,min=10"`
	PrivateToken string `validate:"required,min=10"`
}

func (e AppEnvironment) isProduction() bool {
	return e.Type == production
}

func (e AppEnvironment) isStaging() bool {
	return e.Type == staging
}

func (e AppEnvironment) isLocal() bool {
	return e.Type == local
}
