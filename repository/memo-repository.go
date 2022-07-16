package repository

import (
	"github.com/someday-94/TypeGoMongo-Server/model"
)

type MemoRepository interface {
	Save(memo *model.Memo)
	Update(memo model.Memo)
	Delete(memo model.Memo)
	FindAll() []*model.Memo
}

type memoRepository struct {
	db Database
	databaseName string
	collectionName string
}

func NewMemoRepository(db Database) MemoRepository {
	db.AutoMigrate(&model.Memo{}, &model.User{})

	return &memoRepository{
		db: db,
	}
}

func (memoRepo *memoRepository) Save(memo *model.Memo) {
	memoRepo.db.Create(memo)
}

func (memoRepo *memoRepository) Update(memo model.Memo) {
	memoRepo.db.Save(&memo)
}

func (memoRepo *memoRepository) Delete(memo model.Memo) {
	memoRepo.db.Delete(&memo)
}

func (memoRepo *memoRepository) FindAll() []*model.Memo {
	var memos []model.Memo
	memoRepo.db.FindAll(&memos)

	var results []*model.Memo
	for _, m := range memos {
		results = append(results, &m)
	}

	return results
}
