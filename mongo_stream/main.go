package main

import (
	"context"
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoURI   = flag.String("mongouri", "mongodb://db01.dc1.osh.telmax.ca:27017", "test_database")
	collection = flag.String("CRUD Database", "network", " for example")
	DBClient   *mongo.Client
	AppDB      *mongo.Database
)

func init() {
	flag.Parse()

	DBClient = DBConnect(*MongoURI)
	if DBClient != nil {
		AppDB = DBClient.Database(*collection)
	}
}

func main() {
	cur, err := AppDB.Collection("as_flow").Watch(context.TODO(), mongo.Pipeline{})
	if err != nil {
		panic(err)
	}

	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var data bson.M
		if err := cur.Decode(&data); err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", data)
	}
}

func DBConnect(URI string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Problem connecting MongoURI %v, error : %v ", URI, err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Problem pinging server %v, err %v", URI, err)
	} else {
		log.Infof("Connected to Mongodb server %v", URI)
	}

	return client
}
