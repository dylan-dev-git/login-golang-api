package dbconnector

import (
	"context"
	"fmt"
	"log"

	"github.com/dylan-dev-git/login-microservice-api/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var usersCollection *mongo.Collection
var userinfoCollection *mongo.Collection

var ctx = context.TODO()
var client *mongo.Client

func init() {
	// credential := options.Credential{
	// 	Username:   env.AUTHUSER,
	// 	Password:   env.AUTHPW,
	// 	AuthSource: env.AUTHSOURCE,
	// }
	// clientOptions := options.Client().ApplyURI(env.DBADDR).SetAuth(credential)

	clientOptions := options.Client().ApplyURI(env.DBADDR)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected MongoDB.")

	DB := client.Database(env.DBNAME)

	usersCollection = DB.Collection("users")
	userinfoCollection = DB.Collection("user_info")
}

func CloseDB() {
	client.Disconnect(ctx)
	fmt.Println("Disconnected MongoDB.")
}

func GetUserCollection() *mongo.Collection {
	return usersCollection
}

func GetUserInfoCollection() *mongo.Collection {
	return userinfoCollection
}
