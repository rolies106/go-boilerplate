package services

import (
	"encoding/json"
	"errors"
	"log"
	"mortred/utils"
	"os"
	"time"

	"github.com/streadway/amqp"
)

// GetAmqp Get current connection to RabbitMQ
func GetAmqp(topicName string) (*amqp.Channel, *amqp.Connection) {

	connection, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		utils.Log("error", err)
	}

	channel, err := connection.Channel()
	if err != nil {
		utils.Log("error", err)
	}

	err = channel.ExchangeDeclare(topicName, "topic", true, false, false, false, nil)
	if err != nil {
		utils.Log("error", err)
	}

	return channel, connection
}

func SendToQueue(topicName string, keyRouting string, m interface{}) (err error) {

	channel, connection := GetAmqp(topicName)

	bodyMsg, _ := json.Marshal(m)
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
		ContentType:  "application/json",
		Body:         []byte(bodyMsg),
	}

	err = channel.Publish(topicName, keyRouting, false, false, msg)
	if err != nil {
		// Since publish is asynchronous this can happen if the network connection
		// is reset or if the server has run out of resources.
		log.Fatalf("basic.publish: %v", err)
	}

	returnErrChan := channel.NotifyReturn(make(chan amqp.Return))

	go func() {
		notification, ok := <-returnErrChan
		if !ok {
			// Channel was closed.
		}

		utils.Log("error", errors.New(notification.ReplyText))
	}()

	_, err = channel.QueueDeclare(topicName, true, false, false, false, nil)

	if err != nil {
		// Since publish is asynchronous this can happen if the network connection
		// is reset or if the server has run out of resources.
		utils.Log("error", err)
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind(topicName, "#", topicName, false, nil)

	if err != nil {
		// Since publish is asynchronous this can happen if the network connection
		// is reset or if the server has run out of resources.
		utils.Log("error", err)
	}

	defer connection.Close()

	return
}
