package lib

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	*gorm.DB
}

func NewPostgresDB(env Env, logger Logger) PostgresDB {
	host := env.DBHost
	port := env.DBPort
	user := env.DBUsername
	password := env.DBPassword
	dbname := env.DBName

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.GetGormLogger(),
	})

	if err != nil {
		logger.Info("dsn:", dsn)
		logger.Panic(err)
	}

	logger.Info("PostgreSQL database connection established")

	return PostgresDB{
		DB: db,
	}
}
