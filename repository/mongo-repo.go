package repository

import (
	"context"
	"fmt"
	"log"

	//"os"
	"time"

	"github.com/someday-94/TypeGoMongo-Server/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepo interface {
	Create(memo *model.Memo)
	FindAll() []*model.Memo

}

type mongoRepo struct {
	client *mongo.Client
}

const (
	DATABASE = "graphql"
	COLLECTION = "memos"
)

func NewMongoRepo() MongoRepo {

	//mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	//MONGODB := os.Getenv("MONGODB")
	MONGODB := "mongodb://root:example@127.0.0.1:27017"

	clientOptions := options.Client().ApplyURI(MONGODB)

	clientOptions = clientOptions.SetMaxPoolSize(50)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)

	dbClient, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = dbClient.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("Connected to MongoDB")

	return &mongoRepo{
		client: dbClient,
	}
}

func (m *mongoRepo) Create(memo *model.Memo) {
	collection := m.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), memo)

	if err != nil {
		log.Fatal(err)
	}
}

func (m *mongoRepo) FindAll() []*model.Memo {
	collection := m.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())

	var result []*model.Memo
	for cursor.Next(context.TODO()) {
		var v *model.Memo
		err := cursor.Decode(&v)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, v)
	}

	return result
}