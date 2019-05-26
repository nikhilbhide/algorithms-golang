package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

//process the message
func process(msg *kafka.Message, err error) {
	//check for errors
	if err == nil {
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	} else {
		// The client will automatically try to recover from all errors.
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	}
}

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "new_group_test",
		"auto.offset.reset": "latest",
		"auto.commit.interval.ms":10,
		})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"newtopic"}, nil)

		//poll to the topic and consume the message
		msg, err := c.ReadMessage(-1)

		//invoke go routine to process message
		go process(msg,err)

	c.Close()
}