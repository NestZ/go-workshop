package repos

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var db *mongo.Database

func Init() *Repos {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	db = client.Database("emailCampaign")

	return &Repos{
		Email: &EmailRepository{
			col: db.Collection("emails"),
		},
		Campaign: &CampaignRepository{
			col: db.Collection("campaigns"),
		},
	}
}

type Repos struct {
	Email    *EmailRepository
	Campaign *CampaignRepository
}
