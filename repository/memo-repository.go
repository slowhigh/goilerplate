package repository

import (
	"log"

	"github.com/someday-94/TypeGoMongo-Server/model"
	"go.mongodb.org/mongo-driver/bson"
)

type MemoRepository interface {
	Save(memo model.Memo)
	Update(memo model.Memo)
	Delete(memo model.Memo)
	FindAll() []*model.Memo
}

type memoRepository struct {
	dbConn MongoDB
	dbName string
	collectionName string
}

func NewMemoRepository(dbConn MongoDB, dbName string, collectionName string) MemoRepository {
	return &memoRepository{
		dbConn: dbConn,
		dbName: dbName,
		collectionName: collectionName,
	}
}

func (memoRepo *memoRepository) Save(memo model.Memo) {
	memoRepo.dbConn.InsertOne(memoRepo.dbName, memoRepo.collectionName, memo)
}

func (memoRepo *memoRepository) Update(memo model.Memo) {
	memoRepo.dbConn.UpdateByID(memoRepo.dbName, memoRepo.collectionName, memo.ID, memo)
}

func (memoRepo *memoRepository) Delete(memo model.Memo) {
	memoRepo.dbConn.DeleteOne(memoRepo.dbName, memoRepo.collectionName, model.Memo{ ID: memo.ID })
}

func (memoRepo *memoRepository) FindAll() []*model.Memo {
	result, err := memoRepo.dbConn.Find(memoRepo.dbName, memoRepo.collectionName, bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	for v := range result {
		memo, err := v.(*model.Memo)
		
	}


	con, err := []*model.Memo(result)

}
