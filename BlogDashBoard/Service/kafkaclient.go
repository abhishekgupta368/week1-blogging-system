package blogdashboardservice

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (dbs *DashBoardService) PushCommentToQueue(topic string, message []byte) error {
	producer := dbs.conn
	// defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}

type RequestBody struct {
	RequestType string      `json:"type"`
	Body        interface{} `json:"body"`
	CallBackUrl string      `json:"callback"`
}
type DeleteBody struct {
	Name string `json:"name"`
}
