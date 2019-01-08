package main

import (
	"flag"
	"log"
	"strings"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	var bindingAll bool
	flag.BoolVar(&bindingAll, "all", false, "绑定全部的 routing")
	flag.Parse()

	routingKeys := []string{"info", "warning", "error"}
	if !bindingAll {
		routingKeys = routingKeys[0:1]
	}
	log.Printf("Bind routing keys are %s\n", strings.Join(routingKeys, "&"))

	conn, err := amqp.Dial("amqp://root:root123@localhost:8089")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 与 producer 中的代码类似，创建 exchange
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
	// 这里创建匿名的 queue，设置 exclusive 为 true
	// 链接断开，queue 删除
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// 将 queue 与 exchange 绑定
	// 设置 routingKey
	for _, routing := range routingKeys {
		err = ch.QueueBind(
			q.Name,        // name
			routing,       // routing key
			"logs_direct", // source name
			false,         // no-wait
			nil,
		)
		failOnError(err, "Failed to bind a queue")
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
