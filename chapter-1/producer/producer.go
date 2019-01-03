package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	const (
		exchange   = "exchange_demo"
		routingKey = "routingkey_demo"
		// queue      = "queue_demo"
	)

	conn, err := amqp.Dial("amqp://root:root123@localhost:8089")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange,
		"direct",
		true,  // durable
		false, // auto-delete
		false, // internal
		false, // no-wait
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	message := "Hello World!"

	err = ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	failOnError(err, "Failed to publish a message")
	log.Printf("[x] Send %s", message)
}
