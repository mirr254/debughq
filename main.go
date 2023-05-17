package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mirr254/debughq/initializers"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load Env vars")
	}

	//connect db
	initializers.ConnectDB(&config)
	server = gin.Default()

}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load Env vars")
	}

	router := server.Group("/api/v1")
	router.GET("/healthchecker", func( ctx *gin.Context ) {
		message := "Welcome to DEBUGHQ"
		ctx.JSON(http.StatusOK, gin.H{"Status": "Success", "Message": message})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}