package db

import (
	"errors"
	"log"
	"github.com/Shopify/sarama"
	"fmt"
	"os"
	"os/signal"
)

type kafkaObj struct {
	sarama.SyncProducer
	sarama.Consumer
}


var brokers = []string{"127.0.0.1:9092"}
var msgcount = 0

func NewKafka() Store {
	//settings := getCredentials()

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := sarama.NewConsumer(brokers, nil)

	if err != nil {
		panic(err)
	}

	return &kafkaObj{
		producer,
		consumer,
	}
}

func (db *kafkaObj) Add(id string, user User) error {
	msg := &sarama.ProducerMessage{
		Topic:     "users",
		Partition: 99,
		Value:     sarama.StringEncoder(fmt.Sprintf("%s %s %s", user.Name, user.Address, user.Contact)),
	}
	_, _, err := db.SyncProducer.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	return nil
}

func (db *kafkaObj) Get(id string) (User, error) {
	consumer, err := db.Consumer.ConsumePartition("users", 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgcount++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	return User{}, errors.New("not found")
}

func (db *kafkaObj) Update(user User) error {
	return nil
}

func (db *kafkaObj) Delete(id string) error {
	return nil
}

