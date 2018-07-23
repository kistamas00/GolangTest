package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
	"net/http"
	"db"
	"strconv"
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
			c.JSON(http.StatusOK, result)
		}
	})
	router.PUT("/votes/:id/inc/:optionIndex", func(c *gin.Context) {

		id := c.Param("id")
		optionIndex, err := strconv.Atoi(c.Param("optionIndex"))

		if err == nil {

			success := db.IncreaseVoteCount(id, optionIndex)

			if success {
				c.JSON(http.StatusOK, nil)
			} else {
				c.JSON(http.StatusInternalServerError, nil)
			}

		} else {
			c.JSON(http.StatusBadRequest, nil)
		}
	})
}