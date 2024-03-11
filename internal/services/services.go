package services

import (
	"fmt"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

type Value struct {
	Value int
	mutex sync.RWMutex
}

func NewValue() *Value {
	return &Value{
		Value: 0,
		mutex: sync.RWMutex{},
	}
}

func (v *Value) AddTotal(value int) {
	v.mutex.RLock()
	defer v.mutex.RUnlock()
	v.Value += value
}

func Producer(value int) {
	//create kafka producer
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalf("error creating Kafka producer: %s", err)
	}

	defer producer.Close()

	//create message
	msgValue := fmt.Sprintf("%d", value)
	message := &sarama.ProducerMessage{
		Topic: "count",
		Value: sarama.StringEncoder(msgValue),
	}
	fmt.Println(message.Value)

	//send message

	_, _, err = producer.SendMessage(message)
	if err != nil {
		log.Fatalf("failed to send kafka message: %s", err)
	}

	log.Print("message sent successfully")
}
