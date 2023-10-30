package queue

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func SendToQueue(data []byte) {
	fmt.Println("Setup connection with RABBIT_MQ")
	conn, err := amqp.Dial(os.Getenv("RABBIT_MQ_CONNECTION_STRING"))
	if err != nil {
		fmt.Println("Cannot create connection to RabbitMQ")
		fmt.Println(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("Cannot create channel in RabbitMQ")
		fmt.Println(err)
		return
	}
	defer ch.Close()

	fmt.Println("SENT DATA TO RABBITMQ")

	err = ch.Publish(
		os.Getenv("RABBIT_MQ_EVENTS_EXCHANGE_NAME"),
		os.Getenv("RABBIT_MQ_EVENTS_ROUTE_KEY"),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
	if err != nil {
		fmt.Println("Cannot publsh message in RabbitMQ")
		fmt.Println(err)
		return
	}

	fmt.Println("DATA is published to RABBITMQ")
}
