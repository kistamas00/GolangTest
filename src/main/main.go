package main

import (
	"db"
	"controller"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
)

func main() {

	db.Init()

	log := logrus.New()
	router := gin.New()

	router.LoadHTMLGlob("./templates/*")

	router.Use(ginlogrus.Logger(log), gin.Recovery())

	controller.AddPublicAP(router)
	controller.AddAdminAP(router)

	router.Run()
}