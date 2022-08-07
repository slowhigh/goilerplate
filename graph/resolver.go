package graph

import "github.com/someday-94/TypeGoMongo-Server/repository"

// This file will not be regenerated automatically.
//go:generate go run github.com/99designs/gqlgen generate
// It serves as dependency injection for your app, add any dependencies you require here.

// go generate ./...
type Resolver struct {
	userRepo *repository.UserRepository
	memoRepo *repository.MemoRepository
}

func NewResolver() *Resolver {
	mongodb := repository.NewMongoDB("127.0.0.1", "27017", "root", "example")

	return &Resolver{
		userRepo: repository.NewUserRepository(mongodb),
		memoRepo: repository.NewMemoRepository(mongodb),
	}
}
