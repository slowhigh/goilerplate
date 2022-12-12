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

// NewUserService initialize user service.
func NewUserService(logger lib.Logger, userRepository repositories.UserRepository) UserService {
	return UserService{
		logger:         logger,
		userRepository: userRepository,
	}
}

// GetUserInfoByID gets the user matching given the id.
func (us UserService) GetUserInfoByID(id string) (user models.User, err error) {
	return us.userRepository.FindOne(models.User{ID: id})
}

// GetUserInfoByEmail gets the user matching given the email.
func (us UserService) GetUserInfoByEmail(email string) (user models.User, err error) {
	return us.userRepository.FindOne(models.User{Email: email})
}

// GetUsers get all users.
func (us UserService) GetUsers() (users []models.User, err error) {
	return us.userRepository.FindAll(models.User{})
}

// CreateUser inserts the user.
func (us UserService) CreateUser(user models.User) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		us.logger.Error(err)
		return err
	}

	user.Password = string(encryptedPassword)

	err = us.userRepository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser updates the user matching the given user. but the user must contain an ID.
func (us UserService) UpdateUser(user models.User) error {
	return us.userRepository.Update(user)
}

// DeleteUser deletes the user matching the given user's ID.
func (us UserService) DeleteUser(id string) error {
	return us.userRepository.Delete(models.User{ID: id})
}

// ExistsEmail verify that at least one user matches the given email.
func (us UserService) ExistsEmail(email string) (bool, error) {
	return us.userRepository.Exists(models.User{Email: email})
}
