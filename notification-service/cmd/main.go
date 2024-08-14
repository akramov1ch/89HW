package main

import (
	"log"
	"notification-service/consumer"
	"notification-service/handlers"
)

func main() {
	consumer := consumer.NewOrderConsumer(consumer.GetKafkaBrokers(), consumer.GetTopic())
	defer consumer.Close()

	go handlers.HandleNotifications(consumer)

	log.Println("Starting notification-service")
	select {} // Service is kept alive
}
