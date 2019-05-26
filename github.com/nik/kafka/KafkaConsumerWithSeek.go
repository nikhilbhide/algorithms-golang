package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	//create at least consumer
	//enable.auto.commit is true and sync processing and auto.commit.interval.ms = 10000 -> atleast once/atmost once
	//enable.auto.commit is true and async processing and auto.commit.interval.ms = 0 -> atmost once
	//enable.auto.commit is false -> atleast once

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servergit s":  "localhost",
		"group.id":           "default_group_specific_parition-2",
		"auto.offset.reset":  "earliest",
	})

	//raise the panic in case of error
	if err != nil {
		panic(err)
	}

	topic0 := "multiple_part_topic"
	partitions := kafka.TopicPartitions{
		{Topic: &topic0, Partition: 2,Offset:10}}

	//assign to specific partition
	err = c.Assign(partitions)

	//check for the error
	if err != nil {
		fmt.Printf("\nAssign failed: %s", err)
	}

	//check the assignment
	assignment, err := c.Assignment()
	if err != nil {
		fmt.Printf("\nAssign failed: %s", err)
	}

	fmt.Printf("\nCompare Assignment %v to original list of partitions %v\n",
		assignment, partitions)

	//seek to a pariticular offset
	topicParition:= kafka.TopicPartition{Topic: &topic0, Partition: 2,Offset:kafka.OffsetEnd}

	err = c.Seek(topicParition,-1)
	if(err==nil) {
		for {
			msg, err := c.ReadMessage(-1)

			//check for errors
			if err == nil {
				fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			} else {
				// The client will automatically try to recover from all errors.
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}
		}
	}

	c.Close()
}
