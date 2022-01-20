package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"host.local/go/golang-todo-api/database"
	"host.local/go/golang-todo-api/handlers"
	"host.local/go/golang-todo-api/middlewares"
)

func main() {
	log.Info("Starting the application")
	database.Init()

	todoHandler := handlers.NewTodo()

	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	router.GET("/get", todoHandler.GetTodos)

	router.POST("/add", todoHandler.CreateTodo)

	router.PUT("/update/:id", todoHandler.UpdateTodo)

	router.DELETE("/delete/:id", todoHandler.DeleteTodo)

	router.Run(":4000")

}
