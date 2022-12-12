package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`

	LogOutput string `mapstructure:"LOG_OUTPUT"`
	LogLevel  string `mapstructure:"LOG_LEVEL"`

	JWTSecret string `mapstructure:"JWT_SECRET"`

	PostgresUserName string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASS"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresDB       string `mapstructure:"POSTGRES_DATABASE"`

	RedisPassword string `mapstructure:"REDIS_PASS"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPort     string `mapstructure:"REDIS_PORT"`
	RedisName     string `mapstructure:"REDIS_NAME"`
}

// create a new environment
func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot read config file")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("environment cannot be loaded: ", err)
	}

	return env
}
