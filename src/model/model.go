package model

import "github.com/mongodb/mongo-go-driver/bson/objectid"

type Vote struct {
	Id objectid.ObjectID
	Question string
	Options []string
	Votes []int64
}