package main

import (
	"context"
	"flag"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoURI       = flag.String("mongouri", "mongodb://localhost:27017", "test database")
	testCollection = flag.String("test", "test", "testing")
	DBClient       *mongo.Client
	AppDB          *mongo.Database
)

func init() {
	flag.Parse()
	DBClient = DBConnect(*MongoURI)
	if DBClient != nil {
		AppDB = DBClient.Database(*testCollection)
	}
}

func DBConnect(URI string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Problem connecting to Mongo URI %v, received errer %v", URI, err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Problem pinging server %v, received error %v", URI, err)
	} else {
		log.Infof("Connected to MongoDB %v", URI)
	}
	return client
}

func main() {
}
