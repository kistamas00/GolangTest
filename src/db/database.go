package db

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"model"
)

const dbName  = "GolangWebAppDB"

var database *mongo.Database

func Init()  {

	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil { log.Fatal(err) }

	err = client.Connect(context.Background())
	if err != nil { log.Fatal(err) }

	database = client.Database(dbName)

	users := database.Collection("users")
	err = users.Drop(context.Background())
	if err != nil { log.Fatal(err) }
	users.InsertOne(context.Background(), map[string]string{
		"username" : "admin",
		"password" : "admin",
	})

	votes := database.Collection("votes")
	err = votes.Drop(context.Background())
	if err != nil { log.Fatal(err) }
	votes.InsertOne(context.Background(), map[string]interface{}{
		"question"	:	"Does it work?",
		"options"	:	[]string{"Yes", "No"},
		"votes"		:	[]int{1, 0},
	})
	votes.InsertOne(context.Background(), map[string]interface{}{
		"question"	:	"2x2=?",
		"options"	:	[]string{":/", "5?", "4"},
		"votes"		:	[]int{1, 2, 3},
	})
}

func GetUsers() map[string]string {

	users := database.Collection("users")
	cur, err := users.Find(context.Background(), nil)
	if err != nil { log.Fatal(err) }

	result := make(map[string]string)

	for cur.Next(context.Background()) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil {
			log.Error(err)
		} else {
			result[elem.Lookup("username").StringValue()] = elem.Lookup("password").StringValue()
		}
	}
	if err := cur.Err(); err != nil {
		log.Error(err)
	}

	cur.Close(context.Background())

	return result
}

func GetVotes() []model.Vote {

	votes := database.Collection("votes")
	cur, err := votes.Find(context.Background(), nil)
	if err != nil { log.Fatal(err) }

	var result []model.Vote

	for cur.Next(context.Background()) {
		elem := bson.NewDocument()
		err := cur.Decode(elem)
		if err != nil {
			log.Error(err)
		} else {

			var options []string
			optionsArray := elem.Lookup("options").MutableArray()
			for i := 0; i < optionsArray.Len(); i++ {
				option, err := optionsArray.Lookup(uint(i))
				if err != nil { log.Error(err) }
				options = append(options, option.StringValue())
			}

			var votes []int64
			votesArray := elem.Lookup("votes").MutableArray()
			for i := 0; i < votesArray.Len(); i++ {
				vote, err := votesArray.Lookup(uint(i))
				if err != nil { log.Error(err) }
				votes = append(votes, vote.Int64())
			}

			result = append(result, model.Vote{
				Id			:	elem.Lookup("_id").ObjectID(),
				Question	:	elem.Lookup("question").StringValue(),
				Options		:	options,
				Votes		:	votes,
			})
		}
	}
	if err := cur.Err(); err != nil {
		log.Error(err)
	}

	cur.Close(context.Background())

	return result
}