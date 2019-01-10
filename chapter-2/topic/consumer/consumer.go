package main

import (
	"flag"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var multi bool
	flag.BoolVar(&multi, "multi", false, "绑定多个的 topic")
	flag.Parse()

	topics := []string{"*.orange.*", "*.*.rabbit", "lazy.#"}

	var bindingTopics []string
	if multi {
		bindingTopics = topics[1:]
	} else {
		bindingTopics = topics[0:1]
	}

	conn, err := amqp.Dial("amqp://root:root123@localhost:8089")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // delete when unused
		false,        // internal
		false,        // no-wait
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 绑定 topic
	for _, topic := range bindingTopics {
		err = ch.QueueBind(
			q.Name,       // name
			topic,        // routing
			"logs_topic", // exchange
			false,        // no-wait
			nil,
		)
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
