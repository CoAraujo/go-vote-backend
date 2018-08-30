package rabbitmq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/coaraujo/go-mongo-rabbitmq/domain"
	"github.com/streadway/amqp"
)

var connection *amqp.Connection
var channel *amqp.Channel
var queueName string

func Initializer() {
	connect()
	getChannel()
	getQueueChannel()
}

func CloseConnection() {
	connection.Close()
	channel.Close()
}

func connect() {
	fmt.Println("[RABBITMQ] Connecting...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	failOnError(err, "[RABBITMQ] Failed to connect to RabbitMQ")
	connection = conn

	fmt.Println("[RABBITMQ] Connected sucessfully.")
}

func getChannel() {
	fmt.Println("[RABBITMQ] Setting channel...")

	ch, err := connection.Channel()
	failOnError(err, "[RABBITMQ] Failed to open a channel")
	channel = ch

	fmt.Println("[RABBITMQ] Setup channel successfully.")
}

func getQueueChannel() {
	q, err := channel.QueueDeclare(
		"vote.groupBBB", // name
		true,            // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	failOnError(err, "[RABBITMQ] Failed to declare a queue.")
	queueName = q.Name
}

func SendVote(vote domain.Vote) {
	fmt.Println("[RABBITMQ] Sending vote...")

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(vote)
	reqBodyBytes.Bytes()

	err := channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(reqBodyBytes.Bytes()),
		})
	failOnError(err, "[RABBITMQ] Failed to publish a message")

	fmt.Println("[RABBITMQ] Vote sent sucessfully.")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
