package repository

import (
	"context"
	"log"

	"github.com/someday-94/TypeGoMongo-Server/model"
	"go.mongodb.org/mongo-driver/bson"
)

type MemoRepository interface {
	Save(memo *model.Memo)
	Update(memo *model.Memo)
	DeleteById(id string) *model.Memo
	FindById(id string)
	FindAll() []*model.Memo
}

type memoRepository struct {
	dbConn         MongoDB
	dbName         string
	collectionName string
}

func NewMemoRepository(dbConn MongoDB, dbName string, collectionName string) MemoRepository {
	return &memoRepository{
		dbConn:         dbConn,
		dbName:         dbName,
		collectionName: collectionName,
	}
}

func (memoRepo *memoRepository) Save(memo *model.Memo) {
	memoRepo.dbConn.InsertOne(memoRepo.dbName, memoRepo.collectionName, memo)
}

func (memoRepo *memoRepository) Update(memo *model.Memo) {
	memoRepo.dbConn.UpdateByID(memoRepo.dbName, memoRepo.collectionName, memo.ID, memo)
}

func (memoRepo *memoRepository) DeleteById(id string) {
	result := memoRepo.dbConn.FindOneAndDelete(memoRepo.dbName, memoRepo.collectionName, model.Memo{ID: id})
	{} result.
}

func (memoRepo *memoRepository) FindAll() []*model.Memo {
	memos := make([]*model.Memo, 0)
	result, err := memoRepo.dbConn.Find(memoRepo.dbName, memoRepo.collectionName, bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close(context.TODO())
	if err = result.All(context.TODO(), &memos); err != nil {
		log.Fatal(err)
	}

	return memos
}
