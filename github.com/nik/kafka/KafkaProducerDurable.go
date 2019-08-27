// Example function-based Apache Kafka producer
package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"strconv"
)

func main() {

	//create a producer instance by supplying the bootstrap server details
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"compression.type":   "snappy",
		"compression.level": 0,
		"message.timeout.ms":30000,
		"request.timeout.ms": 15000,
		"enable.idempotence":"true",
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	//name of the topic.
	//here we will have multiple partitions
	topic := "newtopic"

	fmt.Printf("Created Producer %v\n", p)

	// Optional delivery channel, if not specified the Producer object's
	// .Events channel is used.
	deliveryChan := make(chan kafka.Event)
	//push messages to the topic
	for counter := 0; counter < 100; counter ++ {
		key := strconv.Itoa(counter + 1)
		value := "Compressed_Message_No_" + strconv.Itoa(counter+1)

		//push a messaged onto the topic
		//we are adding the key this time around
		err = p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(value),
			Key:            []byte(key),
		}, deliveryChan)

		report := <-deliveryChan
		fmt.Println("received message", report)
		m := report.(*kafka.Message)
		//If there are no errors
		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message counter - %d to topic %s [%d] at offset %v\n",
				counter+1, *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
	}
	//close the channel
	close(deliveryChan)
}
