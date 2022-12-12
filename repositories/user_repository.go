package repositories

import (
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db     *gorm.DB
	logger lib.Logger
}

// NewUserRepository initialize user repository
func NewUserRepository(db lib.Database, logger lib.Logger) UserRepository {
	db.Postgres.AutoMigrate(&models.User{})

	return UserRepository{
		db:     db.Postgres,
		logger: logger,
	}
}

// Create inserts the given user.
func (ur UserRepository) Create(user models.User) error {
	err := ur.db.Create(&user).Error
	if err != nil {
		ur.logger.Error(err)
	}

	return err
}

// FindAll finds all users matching given conditions.
func (ur UserRepository) FindAll(conditions models.User) ([]models.User, error) {
	var users []models.User

	err := ur.db.Where(&conditions).Find(&users).Error
	if err != nil {
		ur.logger.Error(err)
	}

	return users, err
}

// FindOne finds the first user ordered by ID, matching given conditions.
func (ur UserRepository) FindOne(conditions models.User) (models.User, error) {
	var user models.User

	err := ur.db.Where(&conditions).First(&user).Error
	if err != nil {
		ur.logger.Error(err)
	}

	return user, err
}

// Update updates the user matching the given conditions. but the conditions must contain an ID.
func (ur UserRepository) Update(conditions models.User) error {
	tx := ur.db.Model(&models.User{ID: conditions.ID}).Updates(conditions)
	err := tx.Error
	if err != nil {
		ur.logger.Error(err)
	}
	if tx.RowsAffected <= 0 {
		err = gorm.ErrRecordNotFound
		ur.logger.Error(err)
	}

	return err
}

// Delete deletes the user matching the given conditions. but the conditions must contain an ID.
func (ur UserRepository) Delete(conditions models.User) error {
	tx := ur.db.Delete(&conditions)
	err := tx.Error
	if err != nil {
		ur.logger.Error(err)
	}
	if tx.RowsAffected <= 0 {
		err = gorm.ErrRecordNotFound
		ur.logger.Error(err)
	}

	return err
}

// Exists verify that at least one user matches the given conditions.
func (ur UserRepository) Exists(conditions models.User) (bool, error) {
	count := int64(0)

	err := ur.db.Model(&models.User{}).Where(&conditions).Count(&count).Error
	if err != nil {
		ur.logger.Error(err)
		return false, err
	}
	if count <= 0 {
		return false, nil
	}

	return true, nil
}
