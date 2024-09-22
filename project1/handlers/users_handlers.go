package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// Handler for getting a user
func GetUser(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{
        "message": "GET user with ID " + id,
    })
}