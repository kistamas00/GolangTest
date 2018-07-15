package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"db"
)

func AddAdminAP(router *gin.Engine)  {

	authorized := router.Group("/admin", gin.BasicAuth(db.GetUsers()))

	authorized.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"admin.html", db.GetVotes())
	})
}