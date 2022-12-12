package repositories

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/oxyrinchus/goilerplate/lib"
)

var (
	ctx = context.Background()
)

type AuthRepository struct {
	db     lib.Database
	logger lib.Logger
}

// NewAuthRepository initialize auth repository
func NewAuthRepository(db lib.Database, logger lib.Logger) AuthRepository {
	return AuthRepository{
		db:     db,
		logger: logger,
	}
}

// Set Redis `SET key value [expiration]` command.
func (ar AuthRepository) Set(key, value string, expiration time.Duration) error {
	err := ar.db.Redis.Set(ctx, key, value, 0).Err()
	if err != nil {
		ar.logger.Error(err)
	}

	return err
}

// Get Redis `GET key` command. It returns redis.Nil error when key does not exist.
func (ar AuthRepository) Get(key string) (string, error) {
	value, err := ar.db.Redis.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
			ar.logger.Error(err)
			return "", err	
		}

		return "", nil	
	}

	return value, nil
}
