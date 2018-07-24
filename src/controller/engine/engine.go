package engine

import (
	"db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"
	"controller"
)

func NewEngine(db db.DataBase) *gin.Engine {

	log := logrus.New()
	engine := gin.New()

	engine.LoadHTMLGlob("./templates/*")

	engine.Use(ginlogrus.Logger(log), gin.Recovery())
	engine.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	controller.AddPublicAP(engine, db)
	controller.AddAdminAP(engine, db)

	return engine
}
