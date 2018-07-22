package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"db"
	"model"
	"controller/common"
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
	authorized.POST("/votes", func(c *gin.Context) {

		question := c.PostForm("question")
		options := c.PostFormArray("options")

		options = common.Filter(options, func(s string) bool {
			return s != ""
		})

		if question != "" && len(options) >= 2 {

			db.AddVote(model.Vote{
				Question:	question,
				Options:	options,
			})

			c.JSON(http.StatusOK, nil)

		} else {
			c.JSON(http.StatusBadRequest, nil)
		}
	})
	authorized.PUT("/votes/:id", func(c *gin.Context) {

		id := c.Param("id")
		question := c.PostForm("question")
		options := c.PostFormArray("options")

		options = common.Filter(options, func(s string) bool {
			return s != ""
		})

		if question != "" && len(options) >= 2 {

			success := db.EditVote(model.Vote{
				Id:			id,
				Question:	question,
				Options:	options,
			})

			if success {
				c.JSON(http.StatusOK, nil)
			} else {
				c.JSON(http.StatusInternalServerError, nil)
			}

		} else {
			c.JSON(http.StatusBadRequest, nil)
		}
	})
	authorized.DELETE("/votes/:id", func(c *gin.Context) {

		id := c.Param("id")

		success := db.RemoveVote(id)

		if success {
			c.JSON(http.StatusOK, nil)
		} else {
			c.JSON(http.StatusInternalServerError, nil)
		}
	})
}