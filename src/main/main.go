package main

import (
	"db"
	"controller/engine"
	"flag"
)

func main() {

	clearDb := flag.Bool("clearDB", false, "to clear and then insert sample data in to the DB")
	flag.Parse()

	db := db.MongoDB{}
	db.Init()
	if *clearDb == true {
		db.ClearDbAndInsertSamples()
	}

	engine := engine.NewEngine(&db)

	engine.Run()
}