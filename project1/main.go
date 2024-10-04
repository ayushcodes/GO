package main

import (
	"project1/mongodb"
	"project1/handlers"
	"project1/router"
)

func main() {
	mongodb.InitMongo()
	handlers.FetchPosts()
	
	r := router.SetupRouter()
	r.Run(":7007")
}