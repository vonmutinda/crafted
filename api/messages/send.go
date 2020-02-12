package messages

import (
	"fmt"

	"github.com/streadway/amqp"
)

// init connection
func connect() *amqp.Connection {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")  
	if err != nil { 
		failOnError(err, "Failed to connect to RabbitMQ")
	}
	return conn
}

// fail on eror
func failOnError(err error, msg string) {

	if err != nil {
	  fmt.Printf("%s: %s", msg, err)
	}
}

// new chan
func newChannel() *amqp.Channel {
 
	conn, err := connect()

	chan, err := conn.Channel()

	if err != nil {
		failOnError(err, "Failed to open a channel")
	}
	return chan
}

// new queue
func newQueue() *amqp.Queue{ 

	ch := newChannel()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	  )
	
	if err != nil { 
		failOnError(err, "Failed to declare a queue")  
	}
	return q
}

// publish to queue
func sendMessage(){ 

	q := newQueue()

	body := "Hello World!"
	err = ch.Publish(
	  "",     // exchange
	  q.Name, // routing key
	  false,  // mandatory
	  false,  // immediate
	  amqp.Publishing {
		ContentType: "text/plain",
		Body:        []byte(body),
	  })

	if err != nil { 
		failOnError(err, "Failed to publish a message")
	}
}
