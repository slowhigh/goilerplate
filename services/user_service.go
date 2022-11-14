package services

import (
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/models"
	"github.com/oxyrinchus/goilerplate/repositories"
)

type UserService struct {
	logger lib.Logger
	repository repositories.UserRepository
}

// NewUserService initialize user service
func NewUserService(logger lib.Logger, repository repositories.UserRepository) UserService {
	return UserService{
		logger: logger,
		repository: repository,
	}
}

func (us UserService) GetUser(id uint) (user models.User, err error) {
	return user, us.repository.Find(&user, id).Error
}

func (us UserService) GetUsers() (users []models.User, err error) {
	return users, us.repository.Find(&users).Error
}

func (us UserService) CreateUser(user models.User) error {
	return us.repository.Create(&user).Error
}

func (us UserService) UpdateUser(user models.User) error {
	return us.repository.Save(&user).Error
}

func (us UserService) DeleteUser(id uint) error {
	return us.repository.Delete(&models.User{}, id).Error
}