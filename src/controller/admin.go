package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"db"
)

func AddAdminAP(router *gin.Engine)  {

	authorized := router.Group("/admin", gin.BasicAuth(db.GetUsers()))

	// HTML
	authorized.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", nil)
	})

	// JSON
	authorized.GET("/votes", func(c *gin.Context) {
		c.JSON(http.StatusOK, db.GetVotes(nil))
	})
}