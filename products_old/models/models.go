package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Product struct {
		Id    bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Name  string        `json:"name"`
		Price string        `json:"price"`
	}
)
