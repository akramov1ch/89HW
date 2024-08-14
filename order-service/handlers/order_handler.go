package handlers

import (
    "encoding/json"
    "net/http"
    "order-service/producer"

    "github.com/gin-gonic/gin"
)

type Order struct {
    ID     string `json:"id"`
    Status string `json:"status"`
}

func CreateOrder(c *gin.Context) {
    var order Order
    if err := c.BindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    producer := producer.NewOrderProducer(producer.GetKafkaBrokers(), producer.GetTopic())
    defer producer.Close()

    orderJSON, _ := json.Marshal(order)
    err := producer.Produce(order.ID, orderJSON)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to produce message"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order received"})
}
