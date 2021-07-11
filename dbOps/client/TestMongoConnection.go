package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"../connection"

	"gopkg.in/gookit/color.v1"
)

//Tweet - struct to map with mongodb documents
type Tweet struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"createdAt"`
	TweetText string             `bson:"text"`
}

//listAllTweets - Get All issues for collection
func GetAllTweets() ([]Tweet, error) {
	//Define filter query for fetching specific document from collection
	filter := bson.D{{}} //bson.D{{}} specifies 'all documents'
	var issues []Tweet
	//Get MongoDB connection using connectionhelper.
	client, err := connection.GetMongoClient()
	if err != nil {
		return issues, err
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(connection.DB).Collection(connection.TABLE)
	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return issues, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		var t Tweet
		err := cur.Decode(&t)
		if err != nil {
			return issues, err
		}
		issues = append(issues, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(issues) == 0 {
		return issues, mongo.ErrNoDocuments
	}
	return issues, nil
}

func main() {
	tweets, _ := GetAllTweets()
	listAllTweets(tweets)
}

func listAllTweets(issues []Tweet) {
	for i, v := range issues {
		color.Green.Printf("%d: %s   %s", i+1, v.ID, v.CreatedAt, v.TweetText)
	}
}
