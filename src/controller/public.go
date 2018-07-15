package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

func AddPublicAP(router *gin.Engine)  {

	router.Use(static.Serve("/public", static.LocalFile("./public", true)))
}