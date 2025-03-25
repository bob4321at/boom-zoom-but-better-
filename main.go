package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var rooms = []Room{}

func RemoveArrayElement[T any](index_to_remove int, slice *[]T) {
	*slice = append((*slice)[:index_to_remove], (*slice)[index_to_remove+1:]...)
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./templates/*/*")
	r.StaticFile("/matchmakingjs", "./templates/match_making/match_making.js")
	r.StaticFile("/mainjs", "./templates/main/mainjs.js")

	r.GET("/game", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{
			"title":   "Game Room",
			"message": "Welcome to the new game!",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "match_making.tmpl", nil)
	})

	r.POST("/makeRoom", makeRoom)
	r.GET("/getRoom", getRoom)
	r.GET("/addUserToRoom", addUserToRoom)
	r.GET("/removeUserFromRoom", removeUserFromRoom)

	r.Run()
}
