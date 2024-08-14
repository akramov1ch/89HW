package handlers

import (
    "encoding/json"
    "log"
    "notification-service/consumer"
)

type Order struct {
    ID     string `json:"id"`
    Status string `json:"status"`
}

func HandleNotifications(consumer *consumer.OrderConsumer) {
    for {
        msg, err := consumer.Consume()
        if err != nil {
            log.Println("Error consuming message:", err)
            continue
        }

        var order Order
        if err := json.Unmarshal(msg.Value, &order); err != nil {
            log.Println("Error unmarshalling message:", err)
            continue
        }

        if order.Status == "new orders" { // Filtrlash
            log.Println("New order received:", order.ID)
            // Yangi buyurtma bo'yicha amallarni bajarish
        }
    }
}
