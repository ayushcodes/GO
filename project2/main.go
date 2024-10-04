package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type profile struct {
	Id int `json:"id"`
	Name string	`json:"name" binding:"required"`
}

var profiles = []profile{
	{
		Id:1,
		Name:"ayush",
	},
	{
		Id:2,
		Name:"garg",
	},
}

func main() {
	r := gin.Default()  // *gin.Engine

	r.GET("/profiles", func(c *gin.Context){
		c.JSON(http.StatusOK, profiles)
	})

	r.GET("/profiles/:id", func(c *gin.Context){
		id := c.Param("id")
		int_id , _ := strconv.Atoi(id)
		for _, profile := range profiles {
			if profile.Id == int_id {
				c.JSON(http.StatusOK, profile)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"Error":"Profile not found"})
	})

	r.POST("/profiles", func(c *gin.Context){
		var newProfile profile

		err := c.ShouldBindJSON(&newProfile)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		profiles = append(profiles, newProfile)
		c.JSON(http.StatusCreated, newProfile)
	})

	r.Run()
}