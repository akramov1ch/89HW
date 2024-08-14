package main

import (
	"log"
	"order-service/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/order", handlers.CreateOrder)

	log.Println("Starting order-service on port 8080")
	r.Run(":8080")
}
