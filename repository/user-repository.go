package repository

import (
	"github.com/someday-94/TypeGoMongo-Server/graph/model"
)

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(user model.User)
	FindAll() []*model.User
}

type userRepository struct {
	db Database
}

func NewUserRepository(db Database) UserRepository {
	db.AutoMigrate(&model.User{})

	return &userRepository{
		db: db,
	}
}

func (userRepo *userRepository) Save(user model.User) {
	userRepo.db.Create(&user)
}

func (userRepo *userRepository) Update(user model.User) {
	userRepo.db.Save(&user)
}

func (userRepo *userRepository) Delete(user model.User) {
	userRepo.db.Delete(&user)
}

func (userRepo *userRepository) FindAll() []*model.User {
	var users []model.User
	userRepo.db.FindAll(&users)

	var results []*model.User
	for _, m := range users {
		results = append(results, &m)
	}

	return results
}
