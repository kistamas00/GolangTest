package main

import (
	"db"
	"controller/engine"
)

func main() {

	db := db.MongoDB{}

	db.Init()

	engine := engine.NewEngine(&db)

	engine.Run()
}