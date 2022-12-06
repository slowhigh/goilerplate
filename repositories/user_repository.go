package repositories

import (
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*gorm.DB
	logger lib.Logger
}

func NewUserRepository(db lib.Database, logger lib.Logger) UserRepository {
	db.Postgres.AutoMigrate(&models.User{})

	return UserRepository{
		DB: db.Postgres,
		logger:   logger,
	}
}
