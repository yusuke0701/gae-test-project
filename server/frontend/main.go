package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("", "./static")

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
}
