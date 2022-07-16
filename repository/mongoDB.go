package repository

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB interface {
	Create(database string, collection string, interface{})
	Read(database string, collection string, interface{})
	Update(database string, collection string, interface{})
	Delete(database string, collection string, interface{})
}

type mongoDB struct {
	client *mongo.Client
}

func NewMongoRepo(host, port, id, pw string) MongoRepo {

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

	return &mongoRepo{
		client: dbClient,
	}
}