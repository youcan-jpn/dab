package env

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap/zapcore"
)

const ServiceName = "dab"

type (
	Environment struct {
		Port int `envconfig:"PORT" default:"1323"`

		DBEngine string `envconfig:"DB_ENGINE" default:"mysql"`
		DBHost   string `envconfig:"DB_HOST" default:"localhost"` // or ip address remotely
		DBPort   int    `envconfig:"DB_PORT" default:"3336"`
		DBUser   string `envconfig:"DB_USER" default:"dab_user"`
		DBPass   string `envconfig:"DB_PASS" default:"passw0rd"`
		DBName   string `envconfig:"DB_NAME" default:"dab_db"`

		LogLevel zapcore.Level `envconfig:"LOG_LEVEL" default:"INFO"`

		DabEnvironment string `envconfig:"EXAM_ENVIRONMENT" default:"development"`
	}
)

const (
	ProdEnv = "production"
	DevEnv  = "development"
)

var (
	env  Environment
	err  error
	once sync.Once
)

func init() {
	_, err := Process()
	if err != nil {
		panic(err)
	}
}

func Process() (Environment, error) {
	once.Do(func() {
		err = envconfig.Process("", &env)
	})
	return env, err
}
