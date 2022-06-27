package main

import (
	"fmt"
	"log"

	"github.com/vinivinci/simulatorDelivery/infra/kafka"

	"github.com/joho/godotenv"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	kafkaInput "github.com/vinivinci/simulatorDelivery/application/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkaInput.Produce(msg)
	}
}
