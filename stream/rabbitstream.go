package stream

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	config "github.com/coaraujo/go-vote-backend/config/rabbit"
	domain "github.com/coaraujo/go-vote-backend/domain"
	"github.com/streadway/amqp"
)

type RabbitStream struct {
	RabbitMQ   *config.RabbitMQ
	Queue      amqp.Queue
	QueueGroup string
}

func NewRabbitStream(conn *config.RabbitMQ, queue string) *RabbitStream {
	q := "vote.groupBBB"
	if queue != "" {
		q = queue
	}
	rabbitStream := RabbitStream{RabbitMQ: conn, QueueGroup: q}
	rabbitStream.Queue = conn.CreateQueue(q)
	return &rabbitStream
}

func (r *RabbitStream) SendVote(vote domain.Vote) {
	fmt.Println("[RABBITMQ] Sending vote...")

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(vote)
	reqBodyBytes.Bytes()

	err := r.RabbitMQ.GetChannel().Publish(
		"",           // exchange
		r.QueueGroup, // routing key
		// "teste.testando", // routing key
		false, // mandatory
		false, // immediate
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
