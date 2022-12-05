package main

import (
	"fmt"
)

func main() {
	var (
		topic  string   = "blogs"
		broker []string = []string{"localhost:9092"}
	)
	fmt.Println("Consumer started")
	blogConsumerService := BlogConsumerService{}
	consumer := NewConsumer(topic, broker, &blogConsumerService)
	consumer.Init()
}
