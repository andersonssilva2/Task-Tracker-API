package handlers

import (
    "net/http"
    "tasktracker/db"
    "tasktracker/models"
    "github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
    if db.DB == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
        return
    }

    rows, err := db.DB.Query("SELECT id, title, completed FROM tasks")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    if rows == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No rows returned"})
        return
    }
    defer rows.Close()

    var tasks []models.Task
    for rows.Next() {
        var t models.Task
        if err := rows.Scan(&t.ID, &t.Title, &t.Completed); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        tasks = append(tasks, t)
    }

    c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
    if db.DB == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is nil"})
        return
    }

    var t models.Task
    if err := c.ShouldBindJSON(&t); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := db.DB.QueryRow(
        "INSERT INTO tasks (title, completed) VALUES ($1, $2) RETURNING id",
        t.Title, t.Completed,
    ).Scan(&t.ID)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, t)
}
