package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"net/http"
	"db"
)

func AddPublicAP(router *gin.Engine)  {

	// STATIC
	router.Use(static.Serve("/public", static.LocalFile("./public", true)))

	//HTML
	router.GET("/letsVote/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "vote.html", nil)
	})

	// JSON
	router.GET("/votes/:id", func(c *gin.Context) {

		id := c.Param("id")
		result := db.GetVotes([]string{id})

		if result == nil {
			c.JSON(http.StatusNotFound, nil)
		} else {
			c.JSON(http.StatusOK, db.GetVotes([]string{id}))
		}
	})
}