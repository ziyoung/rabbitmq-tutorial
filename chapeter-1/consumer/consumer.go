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
		queue      = "queue_demo"
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
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	_, err = ch.QueueDeclare(
		queue, // name
		false, // durable
		false, // delete when used
		true,  // exclusive
		false, // no-wait
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		queue,      // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,      // no-wait
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		queue, // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-wait
		false, // no-local
		nil,   // args
	)
	failOnError(err, "Failed to regester a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
