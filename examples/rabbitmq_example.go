package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// type Vote struct {
// 	Option    int    `json:"option"`
// 	ParedaoID string `json:"paredaoId"`
// }

func rabbit() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"vote.groupBBB", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	failError(err, "Failed to declare a queue")

	vote := Vote{1, "1"}
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(vote)
	reqBodyBytes.Bytes()

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(reqBodyBytes.Bytes()),
		})
	failError(err, "Failed to publish a message")
}

func failError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
