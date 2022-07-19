package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB interface {
	InsertOne(dbName string, collectionName string, document interface{}) (*mongo.InsertOneResult, error)
	InsertMany(dbName string, collectionName string, documents []interface{}) (*mongo.InsertManyResult, error)
	FindOne(dbName string, collectionName string, filter interface{}) *mongo.SingleResult
	Find(dbName string, collectionName string, filter interface{}) ([]*interface{}, error)
	UpdateByID(dbName string, collectionName string, id interface{}, update interface{}) (*mongo.UpdateResult, error)
	UpdateOne(dbName string, collectionName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	UpdateMany(dbName string, collectionName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	DeleteOne(dbName string, collectionName string, filter interface{}) (*mongo.DeleteResult, error)
	DeleteMany(dbName string, collectionName string, filter interface{}) (*mongo.DeleteResult, error)
}

type mongoDB struct {
	client *mongo.Client
}

func NewMongoDB(host, port, id, pw string) MongoDB {

	//mongodb+srv://USERNAME:PASSWORD@HOST:PORT
	//MONGODB := os.Getenv("MONGODB")
	MONGODB := fmt.Sprintf("mongodb://%s:%s@%s:%s", id, pw, host, port)

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

	return &mongoDB{
		client: dbClient,
	}
}

func (db *mongoDB) InsertOne(dbName string, collectionName string, document interface{}) ([]*interface{}, error) {
	collection := db.client.Database(dbName).Collection(collectionName)
	//return collection.InsertOne(context.TODO(), document)

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())

	var result []*interface{}
	for cursor.Next(context.TODO()) {
		var v *interface{}
		err := cursor.Decode(&v)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, v)
	}

	return result, nil
}

func (db *mongoDB) InsertMany(dbName string, collectionName string, documents []interface{}) (*mongo.InsertManyResult, error) {
	collection := db.client.Database(dbName).Collection(collectionName)
	return collection.InsertMany(context.TODO(), documents)
}

func (db *mongoDB) FindOne(dbName string, collectionName string, filter interface{}) *mongo.SingleResult {
	collection := db.client.Database(dbName).Collection(collectionName)
	return collection.FindOne(context.TODO(), filter)
}

func (db *mongoDB) Find(dbName string, collectionName string, filter interface{}) (*mongo.Cursor, error) {
	collection := db.client.Database(dbName).Collection(collectionName)
	return collection.Find(context.TODO(), filter)
}

func (db *mongoDB) UpdateByID(dbName string, collectionName string, id interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := db.client.Database(dbName).Collection(collectionName)
	return collection.UpdateByID(context.TODO(), id, update)
}

func (db *mongoDB) UpdateOne(dbName string, collectionName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := db.client.Database(dbName).Collection(collectionName)
	return collection.UpdateOne(context.TODO(), filter, update)
}

func (db *mongoDB) UpdateMany(dbName string, collectionName string, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := db.client.Database(dbName).Collection(collectionName)
	return collection.UpdateMany(context.TODO(), filter, update)
}

func (db *mongoDB) DeleteOne(dbName string, collectionName string, filter interface{}) (*mongo.DeleteResult, error) {
	collection := db.client.Database(dbName).Collection(collectionName)
	return collection.DeleteOne(context.TODO(), filter)
}

func (db *mongoDB) DeleteMany(dbName string, collectionName string, filter interface{}) (*mongo.DeleteResult, error) {
	collection := db.client.Database(dbName).Collection(collectionName)
	return collection.DeleteMany(context.TODO(), filter)
}
