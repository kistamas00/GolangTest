package db

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
	"context"
	"github.com/mongodb/mongo-go-driver/bson"
	"model"
	"controller/common"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"strconv"
)

const mongoDbName  = "GolangWebAppDB"

type MongoDB struct {
	database *mongo.Database
}

func (db *MongoDB) Init()  {

	log.Info("Database Init")

	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil { log.Fatal(err) }

	err = client.Connect(context.Background())
	if err != nil { log.Fatal(err) }

	db.database = client.Database(mongoDbName)
	database := db.database

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
		"question"	:	"First test question",
		"options"	:	[]string{"First option with 1 vote", "Second option with 0 vote"},
		"votes"		:	[]int{1, 0},
	})
	votes.InsertOne(context.Background(), map[string]interface{}{
		"question"	:	"Second test question",
		"options"	:	[]string{"First option with 0 vote", "Second option with 1 vote", "Third option with 8 vote"},
		"votes"		:	[]int{0, 1, 8},
	})
}

func (db *MongoDB) GetUsers() map[string]string {

	log.Info("Database GetUsers")

	database := db.database

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

func (db *MongoDB) GetVotes(filterIds []string) []model.Vote {

	log.WithField("filters", filterIds).Info("Database GetVotes")

	database := db.database

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

			vote := model.Vote{
				Id			:	elem.Lookup("_id").ObjectID().Hex(),
				Question	:	elem.Lookup("question").StringValue(),
				Options		:	options,
				Votes		:	votes,
			}

			if filterIds == nil || common.Include(filterIds, vote.Id) {
				result = append(result, vote)
			}
		}
	}
	if err := cur.Err(); err != nil {
		log.Error(err)
	}

	cur.Close(context.Background())

	return result
}

func (db *MongoDB) AddVote(vote model.Vote) {

	log.WithField("vote", vote).Info("Database AddVote")

	database := db.database

	votes := database.Collection("votes")
	votes.InsertOne(context.Background(), map[string]interface{}{
		"question"	:	vote.Question,
		"options"	:	vote.Options,
		"votes"		:	make([]int, len(vote.Options)),
	})
}

func (db *MongoDB) EditVote(vote model.Vote) bool {

	log.WithField("vote", vote).Info("Database EditVote")

	database := db.database

	objectId, err := objectid.FromHex(vote.Id)

	if err == nil {
		votes := database.Collection("votes")
		votes.UpdateOne(context.Background(), map[string]interface{}{"_id" : objectId}, map[string]interface{}{
			"$set":	map[string]interface{}{
				"question": vote.Question,
				"options":  vote.Options,
				"votes":    make([]int, len(vote.Options)),
			},
		})
		return true
	} else {
		log.WithField("error", err).Error("Error creating objectID from hex: ", vote.Id)
		return false
	}
}

func (db *MongoDB) RemoveVote(id string) bool {

	log.WithField("id", id).Info("Database RemoveVote")

	database := db.database

	objectId, err := objectid.FromHex(id)

	if err == nil {
		votes := database.Collection("votes")
		votes.DeleteOne(context.Background(), map[string]interface{}{"_id" : objectId})
		return true
	} else {
		log.WithField("error", err).Error("Error creating objectID from hex: ", id)
		return false
	}
}

func (db *MongoDB) IncreaseVoteCount(id string, optionIndex int) bool {

	log.WithFields(log.Fields{"id" : id, "optionIndex" : optionIndex}).Info("Database IncreaseVoteCount")

	database := db.database

	objectId, err := objectid.FromHex(id)

	if err == nil {
		votes := database.Collection("votes")
		votes.UpdateOne(context.Background(), map[string]interface{}{"_id" : objectId}, map[string]interface{}{
			"$inc" : map[string]interface{}{
				"votes." + strconv.Itoa(optionIndex) : 1,
			},
		})
		return true
	} else {
		log.WithField("error", err).Error("Error creating objectID from hex: ", id)
		return false
	}
}