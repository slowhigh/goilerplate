package repository

import (
	"github.com/someday-94/TypeGoMongo-Server/graph/model"
)

type MemoRepository interface {
	Save(memo model.Memo)
	Update(memo model.Memo)
	Delete(memo model.Memo)
	FindAll() []model.Memo
}

type memoRepository struct {
	db *database
}

func NewMemoRepository(db *database) MemoRepository {
	db.AutoMigrate(&model.Memo{}, &model.User{})

	return &memoRepository{
		db: db,
	}
}

func (memoRepo *memoRepository) Save(memo model.Memo) {
	memoRepo.db.conn.Create(&memo)
}

func (memoRepo *memoRepository) Update(memo model.Memo) {
	memoRepo.db.conn.Save(&memo)
}

func (memoRepo *memoRepository) Delete(memo model.Memo) {
	memoRepo.db.conn.Delete(&memo)
}

func (memoRepo *memoRepository) FindAll() []model.Memo {
	var memos []model.Memo
	memoRepo.db.conn.Set("gorm:auto_preload", true).Find(&memos)
	return memos
}
