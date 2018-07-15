package db

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
	"context"
)

const dbName  = "GolangWebAppDB"

func Init()  {

	client, err := mongo.NewClient("mongodb://localhost:27017")
	if err != nil { log.Error(err) }

	err = client.Connect(context.TODO())
	if err != nil { log.Error(err) }

	database := client.Database(dbName)

	users := database.Collection("users")
	err = users.Drop(context.TODO())
	if err != nil { log.Error(err) }
	//users.InsertOne(context.Background(), map[string]int{"foo": 2, "bar": 2})

}