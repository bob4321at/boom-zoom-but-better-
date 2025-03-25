package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Messege struct {
	Username string
	Content  string
}

func sendMessege(c *gin.Context) {
	user := c.Query("user")
	messege := c.Query("messege")

	for ri := 0; ri < len(rooms); ri++ {
		if rooms[ri].Name == c.Query("name") {
			room := &rooms[ri]
			room.Messeges = append(room.Messeges, Messege{Username: user, Content: messege})
			fmt.Println(room)
			return
		}
	}
}

type Room struct {
	Name     string
	Users    int
	Messeges []Messege
}

func makeRoom(c *gin.Context) {
	req, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(req))

	temp_room := Room{}
	json.Unmarshal(req, &temp_room)

	rooms = append(rooms, temp_room)
}

func getRoom(c *gin.Context) {
	for ri := 0; ri < len(rooms); ri++ {
		if rooms[ri].Name == c.Query("name") {
			c.JSON(http.StatusOK, rooms[ri])
			fmt.Println(rooms[ri])
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"Error": "incorrect lobby name"})
}

func addUserToRoom(c *gin.Context) {
	room := &Room{}
	for ri := 0; ri < len(rooms); ri++ {
		if rooms[ri].Name == c.Query("name") {
			room = &rooms[ri]
		}
	}

	room.Users += 1
}
func removeUserFromRoom(c *gin.Context) {
	room := &Room{}
	remove_indicie := 0
	for ri := 0; ri < len(rooms); ri++ {
		if rooms[ri].Name == c.Query("name") {
			room = &rooms[ri]
			remove_indicie = ri
		}
	}

	room.Users -= 1

	if room.Users < 1 {
		RemoveArrayElement(remove_indicie, &rooms)
	}

	fmt.Println(rooms)
}
