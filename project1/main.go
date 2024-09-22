package main

import (
	"project1/handlers"
	"project1/router"
)

func main() {
	handlers.FetchPosts()
	r := router.SetupRouter()
	r.Run(":7007")
}