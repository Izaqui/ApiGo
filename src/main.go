package main

import (
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Door", Completed: false},
	{ID: "2", Item: "Window", Completed: false},
	{ID: "3", Item: "Door Window", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos")
	router.Run("localhost:9009")
}
