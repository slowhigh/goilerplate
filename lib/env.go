package lib

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`
	LogOutput   string `mapstructure:"LOG_OUTPUT"`
	DBUsername  string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASS"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBName      string `mapstructure:"DB_NAME"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
}

// create a new environment
func NewEnv() Env {
	env := Env{}
	viper.SetConfigFile(".env") // to-do : development.env, stage.env, production.env

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
