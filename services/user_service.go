package services

import (
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/models"
	"github.com/oxyrinchus/goilerplate/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	logger         lib.Logger
	userRepository repositories.UserRepository
}

// NewUserService initialize user service
func NewUserService(logger lib.Logger, userRepository repositories.UserRepository) UserService {
	return UserService{
		logger:         logger,
		userRepository: userRepository,
	}
}

func (us UserService) FindUserByID(id string) (user models.User, err error) {
	return user, us.userRepository.Where(&models.User{ID: id}).First(&user).Error
}

func (us UserService) FindUserByEmail(email string) (user models.User, err error) {
	return user, us.userRepository.Where(&models.User{Email: email}).First(&user).Error
}

func (us UserService) FindUsers() (users []models.User, err error) {
	return users, us.userRepository.Find(&users).Error
}

func (us UserService) CreateUser(user models.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	} else {
		user.Password = string(encryptedPassword)
	}

	return us.userRepository.Create(&user).Error
}

func (us UserService) UpdateUser(user models.User) error {
	return us.userRepository.Save(&user).Error
}

func (us UserService) DeleteUser(id uint) error {
	return us.userRepository.Delete(&models.User{}, id).Error
}
