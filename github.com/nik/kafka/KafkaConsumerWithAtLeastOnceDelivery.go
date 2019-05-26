package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"strconv"
)

//process the message
func processAtLeastOnce(consumer *kafka.Consumer, msg *kafka.Message, err error) {
	//check for errors
	if err == nil {
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	} else {
		// The client will automatically try to recover from all errors.
		fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	}

	//we require unique id and we want to demonstrate the exactly once processing
	uniqueId := *(msg.TopicPartition.Topic)+ strconv.Itoa(int(msg.TopicPartition.Partition)) + strconv.Itoa(int(msg.TopicPartition.Offset))
	fmt.Println(uniqueId)

	//processing is done for this consumer group and commit the offset
	consumer.CommitMessage(msg)
}

func main() {
	//create at least consumer
	//enable.auto.commit is true and sync processing and auto.commit.interval.ms = 10000 -> atleast once/atmost once
	//enable.auto.commit is true and async processing and auto.commit.interval.ms = 0 -> atmost once
	//enable.auto.commit is false -> atleast once

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "new_group_test",
		"auto.offset.reset": "latest",
		"enable.auto.commit":"false",
		})

	//raise the panic in case of error
	if err != nil {
		panic(err)
	}

	for {

		c.SubscribeTopics([]string{"newtopic"}, nil)

		//poll to the topic and consume the message
		msg, err := c.ReadMessage(-1)

		//invoke go routine to process message
		go processAtLeastOnce(c, msg, err)
	}
	c.Close()
}