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

type User struct {
	Name string `bson:"name"`
	Age  string `bson:"age"`
}

var (
	MongoURI   = flag.String("mongouri", "mongodb://localhost:27017", "test_database")
	collection = flag.String("CRUD Database", "test", " for example")
	DBClient   *mongo.Client
	AppDB      *mongo.Database
)

func main() {
	flag.Parse()
	DBClient = DBConnect(*MongoURI)
	if DBClient != nil {
		AppDB = DBClient.Database(*collection)
	}

	options := options.Find()

	options.SetSort(bson.D{{"age", -1}})
	options.SetLimit(1)

	cur, err := AppDB.Collection("cr").Find(context.TODO(), bson.M{}, options)

	/*var users []bson.M
	if err = cur.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user["name"])
	}*/

	/*var users []interface{}
	if err = cur.All(context.TODO(), &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user)
	}*/

	for cur.Next(context.TODO()) {
		var user User
		err = cur.Decode(&user)
		if err != nil {

		} else {
			fmt.Println(user)
		}
	}

	//cur.Close(context.TODO())

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
