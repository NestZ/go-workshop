package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Email struct {
	ID        primitive.ObjectID `bson: "_id"`
	FirstName string             `bson: "firstName"`
	LastName  string             `bson: "lastName"`
	Email     string             `bson: "email"`
}
