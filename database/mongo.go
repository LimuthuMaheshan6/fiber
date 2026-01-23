package database

import (

	"log"
	"project/env"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"

)

var Client *mongo.Client

func MongoConnection() {
	// Configure server API
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(env.MONGO_DB).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect( opts)
	if err != nil {
		panic(err)
	}


	log.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// Assign to global variable
	Client = client
}

// Call this when your app shuts down
