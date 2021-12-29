package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id    bson.ObjectId `json:"id" bson:"_id"`
	FName string        `json:"fname" bson:"fname"`
	LName string        `json:"lname" bson:"lname"`
	Email string        `json:"email" bson:"email"`
}

