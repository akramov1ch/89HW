package consumer

import (
    "os"
)

func GetKafkaBrokers() []string {
    brokers := os.Getenv("KAFKA_BROKERS")
    return []string{brokers}
}

func GetTopic() string {
    return os.Getenv("KAFKA_TOPIC")
}
