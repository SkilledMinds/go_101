package connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

// Mongodb connection's configs.
const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "Tweeter_APP"
	TABLE            = "tweets"
)


var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once


//GetMongoClient - Return mongodb connection
func GetMongoClient() (*mongo.Client, error) {
	//will execute only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// let's check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
