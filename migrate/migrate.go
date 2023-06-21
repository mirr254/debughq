package main

import (
	"log"
	"fmt"

	"github.com/mirr254/debughq/models"
	"github.com/mirr254/debughq/initializers"
)

func init() {
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load Env Vars")
	}

	initializers.ConnectDB(&config)

}

func main() {
	initializers.DB.AutoMigrate(
		&models.User{},
		&models.ErrorMessage{},
	    &models.Solution{} )
	fmt.Println("? Migration complete")
}