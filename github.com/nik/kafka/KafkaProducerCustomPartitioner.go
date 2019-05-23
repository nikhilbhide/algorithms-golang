// Example function-based Apache Kafka producer
package main

/**
 * Copyright 2016 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
	"strconv"
)

func getParition(key string) int32 {
	partition, _ := strconv.Atoi(key)
	return int32(partition%3)
}
func main() {

	//create a producer instance by supplying the bootstrap server details
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":  "localhost:9092",
		"compression.type":   "snappy",
		"compression.level": 0,
		"message.timeout.ms":30000,
		"request.timeout.ms": 15000,
		"enable.idempotence":"true",
		"linger.ms":20,
		"queue.buffering.max.kbytes":2097151,
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	//name of the topic.
	//here we will have multiple partitions
	topic := "paritionedtopic"

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
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition:getParition(key)},
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