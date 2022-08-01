package repository

import (
	"context"
	"log"

	"github.com/someday-94/TypeGoMongo-Server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemoRepository interface {
	InsertOne(memo *model.Memo)
	Update(memo *model.Memo)
	DeleteById(id string) *model.Memo
	FindAll() []*model.Memo
}

type memoRepository struct {
	collection *mongo.Collection
}

func NewMemoRepository(dbConn *MongoDB, dbName string, collectionName string) MemoRepository {
	return &memoRepository{
		collection: dbConn.client.Database(dbName).Collection(collectionName),
	}
}

func (memoRepo *memoRepository) InsertOne(memo *model.Memo) {
	memoRepo.collection.InsertOne(context.TODO(), memo)
}

func (memoRepo *memoRepository) Update(memo *model.Memo) {
	_, err := memoRepo.collection.UpdateByID(context.TODO(), memo.ID, memo)
	if err != nil {
		log.Fatal(err)
	}
}

func (memoRepo *memoRepository) DeleteById(id string) *model.Memo {
	var memo *model.Memo

	result := memoRepo.collection.FindOneAndDelete(context.TODO(), model.Memo{ID: id})
	result.Decode(memo)

	return memo
}

func (memoRepo *memoRepository) FindAll() []*model.Memo {
	memos := make([]*model.Memo, 0)

	result, err := memoRepo.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	defer result.Close(context.TODO())
	if err = result.All(context.TODO(), &memos); err != nil {
		log.Fatal(err)
	}

	return memos
}

func (memoRepo *memoRepository) FindOneById(id string) *model.Memo {
	var memo *model.Memo

	result := memoRepo.collection.FindOne(context.TODO(), model.Memo {
		ID: id,
	})

	result.Decode(memo)

	return memo
}
