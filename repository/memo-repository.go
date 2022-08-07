package repository

import (
	"context"

	"github.com/someday-94/TypeGoMongo-Server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MemoRepository struct {
	collection *mongo.Collection
}

func NewMemoRepository(dbConn *MongoDB) *MemoRepository {
	return &MemoRepository{
		collection: dbConn.client.Database("graphql").Collection("memos"),
	}
}

func (memoRepo *MemoRepository) InsertOne(memo *model.Memo) {
	res, err := memoRepo.collection.InsertOne(context.TODO(), memo)
	if err != nil {
		panic(err)
	}

	if memo.ID != res.InsertedID.(string) {
		panic("Invalid insertedID")
	}
}

func (memoRepo *MemoRepository) UpdateByID(id string, memo *model.Memo) {
	update := bson.M{"$set": memo}

	res, err := memoRepo.collection.UpdateByID(context.TODO(), id, update)
	if err != nil {
		panic(err)
	}

	if res.MatchedCount != 1 || res.ModifiedCount != 1 {
		panic("Not found matched id")
	}
}

func (memoRepo *MemoRepository) DeleteById(id string) *model.Memo {
	var memo *model.Memo

	res := memoRepo.collection.FindOneAndDelete(context.TODO(), bson.M{"_id": id})
	if err := res.Err(); err != nil {
		panic(err)
	}

	if err := res.Decode(&memo); err != nil {
		panic(err)
	}

	return memo
}

func (memoRepo *MemoRepository) FindAll() []*model.Memo {
	memos := make([]*model.Memo, 0)

	result, err := memoRepo.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	defer result.Close(context.TODO())
	if err = result.All(context.TODO(), &memos); err != nil {
		panic(err)
	}

	return memos
}

func (memoRepo *MemoRepository) FindOneById(id string) *model.Memo {
	var memo *model.Memo

	res := memoRepo.collection.FindOne(context.TODO(), bson.M{"_id": id})

	if err := res.Err(); err != nil {
		panic(err)
	}

	if err := res.Decode(&memo); err != nil {
		panic(err)
	}

	return memo
}
