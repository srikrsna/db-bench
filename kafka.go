package db

import (
	"errors"
	"log"
	"github.com/Shopify/sarama"
	"gopkg.in/alecthomas/kingpin.v2"
	"fmt"
	"os/signal"
	"os"
)

type kafkaProducer struct {
	sarama.SyncProducer
}
type kafkaConsumer struct {
	sarama.Consumer
}
var (
	brokerList        = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
	topic             = kingpin.Flag("topic", "Topic name").Default("important").String()
	partition         = kingpin.Flag("partition", "Partition number").Default("0").String()
	offsetType        = kingpin.Flag("offsetType", "Offset Type (OffsetNewest | OffsetOldest)").Default("-1").Int()
	messageCountStart = kingpin.Flag("messageCountStart", "Message counter start from:").Int()
)

func NewKafkaProducer() Store {
	//settings := getCredentials()
	kingpin.Parse()
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(*brokerList, config)
	if err != nil {
		log.Fatal(err)
	}
	return &kafkaProducer{
		producer,
	}
}

func (db *kafkaProducer) Add(id string, user User) error {
	msg := &sarama.ProducerMessage{
		Topic: *topic,
		Value: sarama.StringEncoder(fmt.Sprintf("%d %d %d", user.Name, user.Contact, user.Address)),
	}
	_, _, err := db.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	return nil
}

func (db *kafkaProducer) Get(id string) (User, error) {



	return User{}, errors.New("not found")
}

func (db *kafkaProducer) Update(user User) error {
	return nil
}

func (db *kafkaProducer) Delete(id string) error {
	return nil
}



func NewKafkaConsumer() kafkaConsumer {
	kingpin.Parse()
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	brokers := *brokerList
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}
	return kafkaConsumer{
		master,
	}
}

func listenKafka(kf *kafkaConsumer) {
	consumer, err := kf.ConsumePartition(*topic, 0, sarama.OffsetOldest)
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
				*messageCountStart++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
}