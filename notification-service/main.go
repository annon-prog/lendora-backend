package main

import (
	"log"
	"notification-service/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//load environment variables
	loadFile()

	//set the router as the default one set up by gin
	router := gin.Default()

	//initialize the routes
	route.Routes(router)

	//start serving  the application
	router.Run(":8888")

}

// load environment variables
func loadFile() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Failed to load environment variables")
	}
}
