package repository

import (
	"context"

	"github.com/someday-94/TypeGoMongo-Server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(dbConn *MongoDB) *UserRepository {
	return &UserRepository{
		collection: dbConn.client.Database("graphql").Collection("users"),
	}
}

func (userRepo *UserRepository) InsertOne(user *model.User) {
	res, err := userRepo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}

	if user.ID != res.InsertedID.(string) {
		panic("Invalid insertedID")
	}
}

func (userRepo *UserRepository) UpdateByID(id string, user *model.User) {
	update := bson.M{"$set": user}

	res, err := userRepo.collection.UpdateByID(context.TODO(), id, update)
	if err != nil {
		panic(err)
	}

	if res.MatchedCount != 1 || res.ModifiedCount != 1 {
		panic("Not found matched id")
	}
}

func (userRepo *UserRepository) DeleteById(id string) *model.User {
	var deletedUser *model.User

	res := userRepo.collection.FindOneAndDelete(context.TODO(), bson.M{"_id": id})
	if err := res.Err(); err != nil {
		panic(err)
	}

	if err := res.Decode(&deletedUser); err != nil {
		panic(err)
	}

	return deletedUser
}

func (userRepo *UserRepository) FindAll() []*model.User {
	users := make([]*model.User, 0)

	result, err := userRepo.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	defer result.Close(context.TODO())
	if err = result.All(context.TODO(), &users); err != nil {
		panic(err)
	}

	return users
}

func (userRepo *UserRepository) FindOneById(id string) *model.User {
	var user *model.User

	res := userRepo.collection.FindOne(context.TODO(), bson.M{"_id": id})

	if err := res.Err(); err != nil {
		panic(err)
	}

	if err := res.Decode(&user); err != nil {
		panic(err)
	}

	return user
}

func (userRepo *UserRepository) FindOneByName(name string) *model.User {
	var user *model.User

	res := userRepo.collection.FindOne(context.TODO(), bson.M{"name": name})

	if err := res.Err(); err != nil {
		panic(err)
	}

	if err := res.Decode(&user); err != nil {
		panic(err)
	}

	return user
}
