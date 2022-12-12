package repositories

import (
	"github.com/oxyrinchus/goilerplate/lib"
	"github.com/oxyrinchus/goilerplate/models"
	"gorm.io/gorm"
)

type MemoRepository struct {
	db     *gorm.DB
	logger lib.Logger
}

// NewMemoRepository initialize memo repository
func NewMemoRepository(db lib.Database, logger lib.Logger) MemoRepository {
	db.Postgres.AutoMigrate(&models.Memo{})

	return MemoRepository{
		db:     db.Postgres,
		logger: logger,
	}
}

// Create inserts the given memo.
func (mr MemoRepository) Create(memo models.Memo) error {
	err := mr.db.Create(&memo).Error
	if err != nil {
		mr.logger.Error(err)
	}

	return err
}

// FindAll finds all memos matching given conditions.
func (mr MemoRepository) FindAll(conditions models.Memo) ([]models.Memo, error) {
	var memos []models.Memo

	err := mr.db.Where(&conditions).Find(&memos).Error
	if err != nil {
		mr.logger.Error(err)
	}

	return memos, err
}

// FindOne finds the first memo ordered by ID, matching given conditions.
func (mr MemoRepository) FindOne(conditions models.Memo) (models.Memo, error) {
	var memo models.Memo

	err := mr.db.Where(&conditions).First(&memo).Error
	if err != nil {
		mr.logger.Error(err)
	}

	return memo, err
}

// Update updates the memo matching the given conditions. but the conditions must contain an ID.
func (mr MemoRepository) Update(conditions models.Memo) error {
	tx := mr.db.Model(&models.Memo{ID: conditions.ID}).Updates(conditions)
	err := tx.Error
	if err != nil {
		mr.logger.Error(err)
	}
	if tx.RowsAffected <= 0 {
		err = gorm.ErrRecordNotFound
		mr.logger.Error(err)
	}

	return err
}

// Delete deletes the memo matching the given conditions. but the conditions must contain an ID.
func (mr MemoRepository) Delete(conditions models.Memo) error {
	tx := mr.db.Delete(&conditions)
	err := tx.Error
	if err != nil {
		mr.logger.Error(err)
	}
	if tx.RowsAffected <= 0 {
		err = gorm.ErrRecordNotFound
		mr.logger.Error(err)
	}

	return err
}
