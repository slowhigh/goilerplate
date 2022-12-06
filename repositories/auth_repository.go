package repositories

import (
	"context"
	"time"

	"github.com/oxyrinchus/goilerplate/lib"
)

var (
	ctx = context.Background()
)

type AuthRepository struct {
	db     lib.Database
	logger lib.Logger
}

func NewAuthRepository(db lib.Database, logger lib.Logger) AuthRepository {
	return AuthRepository{
		db:     db,
		logger: logger,
	}
}

func (ar AuthRepository) Set(key, value string, expiration time.Duration) error {
	return ar.db.Redis.Set(ctx, key, value, 0).Err()
}

func (ar AuthRepository) Get(key string) (value string, err error) {
	return ar.db.Redis.Get(ctx, key).Result()
}
