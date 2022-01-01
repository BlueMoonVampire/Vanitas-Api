package db

import (
	"context"
	"syl-api/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var DB_URL string = ""

func GetUser(id int) (*types.User, error) {
	var user types.User
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	Client, _ = mongo.Connect(ctx, options.Client().ApplyURI(DB_URL))
	err := Client.Database("Sylviorus").Collection("Main").FindOne(ctx, bson.M{"user": id}).Decode(&user)
	return &user, err
}
