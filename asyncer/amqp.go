package asyncer

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// AMQPAsyncer implements the Asyncer Interface for AMQP (i.e. RabbitMQ)
type AMQPAsyncer struct{}

// Name returns name of Asyncer
func (a AMQPAsyncer) Name() string {
	return "AMQP / RabbitMQ"
}

// CallAsync implements Asyncer interface for AWS Lambda
func (a AMQPAsyncer) CallAsync(payload []byte) error {

	log.Printf(
		"ASYNC ENV: %s; PAYLOAD: %s",
		a.Name(),
		payload,
	)

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/text",
			Body:        payload,
		})
	failOnError(err, "Failed to publish a message")

	return nil
}
