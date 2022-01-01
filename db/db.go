package db

import (
	"context"
	"syl-api/types"
	"time"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

var Client *mongo.Collection

var DB_URL string = os.Getenv("DB_URL")

func GetUser(id int) (*types.User, error) {
	var Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	var user types.User

	err := Client.FindOne(Ctx, bson.M{"user": id}).Decode(&user)
	return &user, err
}

func Database() {
	var Ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	k, _ := mongo.Connect(Ctx, options.Client().ApplyURI(DB_URL))
	Client = k.Database("Sylviorus").Collection("Main")

}
