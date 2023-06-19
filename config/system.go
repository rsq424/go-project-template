package config

const (
	SystemEnv   = "NENLY_DEPLOY_ENV"
	SwaggerMode = "NENLY_SWAGGER_MODE"
)

type System struct {
	SystemEnv   string `mapstructure:"NENLY_DEPLOY_ENV"`
	SwaggerMode string `mapstructure:"NENLY_SWAGGER_MODE"`
}
