package router

import (
	"project1/handlers"

	"github.com/gin-gonic/gin"
)

func registerUserRouters(r *gin.Engine) {
	userRouters := r.Group("/user")

	userRouters.GET("/", handlers.GetUser)
}