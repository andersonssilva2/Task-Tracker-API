package main

import (
    "tasktracker/db"
    "tasktracker/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    db.Connect()

    router.GET("/tasks", handlers.GetTasks)
    router.POST("/tasks", handlers.CreateTask)

    router.Run(":8080")
}
