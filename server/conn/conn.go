package conn

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Mongoclient *mongo.Client

//Function to init Connection for database
func InitConn() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://m001-student:m001-mongodb-basics@sandbox-shard-00-02-kjj8q.mongodb.net:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=true"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	Mongoclient = client
}

//Function to return a collection
func Collection() *mongo.Collection {
	return Mongoclient.Database("").Collection("")
}
