package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var repeatTimes int
	flag.IntVar(&repeatTimes, "repeat", 2, "重复次数")
	flag.Parse()
	if repeatTimes < 2 {
		log.Fatal("repeat 次数不能小于 2")
	}

	conn, err := amqp.Dial("amqp://root:root123@localhost:8089")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_direct", // name
		"direct",      // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,
	)
	failOnError(err, "Failed to declare an exchange")

	prefix := "Hello World!"

	for index := 0; index < repeatTimes; index++ {
		routingKey := "info"
		if index%3 == 1 {
			routingKey = "warning"
		} else if index%3 == 2 {
			routingKey = "error"
		}
		message := prefix + " and type is " + routingKey + " id is " + strconv.Itoa(index)
		err = ch.Publish(
			"logs_direct", // exchange
			routingKey,    // routing key
			false,         // mandatory
			false,         // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		failOnError(err, "Failed to publish a message")
		log.Printf("[x] Send %s", message)
	}
}
