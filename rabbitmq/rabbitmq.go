package rabbitmq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/coaraujo/go-mongo-rabbitmq/domain"
	"github.com/streadway/amqp"
)

const (
	rabbitURL  = "amqp://guest:guest@localhost:5672"
	queueGroup = "vote.groupBBB"
)

type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
}

func NewConnection() *RabbitMQ {
	rabbitmq := RabbitMQ{}

	rabbitmq.getConnection()
	rabbitmq.setChannel()
	rabbitmq.setQueueChannel()

	return &rabbitmq
}

func (r *RabbitMQ) CloseConnection() {
	r.connection.Close()
	r.channel.Close()
}

func (r *RabbitMQ) getConnection() {
	fmt.Println("[RABBITMQ] Connecting...")

	conn, err := amqp.Dial(rabbitURL)
	failOnError(err, "[RABBITMQ] Failed to connect to RabbitMQ")
	r.connection = conn

	fmt.Println("[RABBITMQ] Connected sucessfully.")
}

func (r *RabbitMQ) setChannel() {
	fmt.Println("[RABBITMQ] Setting channel...")

	ch, err := r.connection.Channel()
	failOnError(err, "[RABBITMQ] Failed to open a channel")
	r.channel = ch

	fmt.Println("[RABBITMQ] Setup channel successfully.")
}

func (r *RabbitMQ) setQueueChannel() {
	q, err := r.channel.QueueDeclare(
		queueGroup, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "[RABBITMQ] Failed to declare a queue.")
	r.queue = q
}

func (r *RabbitMQ) SendVote(vote domain.Vote) {
	fmt.Println("[RABBITMQ] Sending vote...")

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(vote)
	reqBodyBytes.Bytes()

	err := r.channel.Publish(
		"",         // exchange
		queueGroup, // routing key
		false,      // mandatory
		false,      // immediate
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
