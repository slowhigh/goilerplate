package lib

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Postgres *gorm.DB
	Redis    *redis.Client
}

func NewDatabase(env Env, logger Logger) Database {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", env.PostgresHost, env.PostgresPort, env.PostgresUserName, env.PostgresPassword, env.PostgresDB)
	postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})
	if err != nil {
		logger.Info("dsn:", dsn)
		logger.Panic(err)
	}
	logger.Info("PostgreSQL database connection established")

	redisURL := fmt.Sprintf("redis://:%s@%s:%s/%s", env.RedisPassword, env.RedisHost, env.RedisPort, env.RedisName)
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		logger.Info("redisURL:", redisURL)
		logger.Panic(err)
	}

	redis := redis.NewClient(opt)
	logger.Info("Redis database connection established")

	return Database{
		Postgres: postgres,
		Redis:    redis,
	}
}
