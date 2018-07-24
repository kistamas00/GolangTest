package main

import (
	"db"
	"controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
)

func GetMainEngine() *gin.Engine {

	log := logrus.New()
	engine := gin.New()

	engine.LoadHTMLGlob("./templates/*")

	engine.Use(ginlogrus.Logger(log), gin.Recovery())

	controller.AddPublicAP(engine)
	controller.AddAdminAP(engine)

	return engine
}

func main() {

	db.Init()

	engine := GetMainEngine()

	engine.Run()
}