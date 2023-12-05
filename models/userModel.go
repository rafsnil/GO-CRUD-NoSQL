package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id     bson.ObjectId `json:"Id" bson:"_id"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"Gender" bson:"Gender"`
	Age    int           `json:"age" bson:"age"`
}
