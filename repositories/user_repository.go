package repositories

import (
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/models"
)

type UserRepository struct {
	lib.PostgresDB
	logger lib.Logger
}

func NewUserRepository(db lib.PostgresDB, logger lib.Logger) UserRepository {
	db.AutoMigrate(&models.User{})

	return UserRepository{
		PostgresDB: db,
		logger:     logger,
	}
}
