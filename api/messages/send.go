package messages

import (
	"fmt"

	"github.com/streadway/amqp"
	"github.com/vonmutinda/crafted/api/log" 
)


// init Connection
func Connect() *amqp.Connection {  

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")  
	if err != nil { 
		FailOnError(err, "Failed to connect to RabbitMQ") 
	}
	return conn
}

// fail on eror
func FailOnError(err error, msg string) { 
	log.GetLogger().Info("%s: %s", msg, err)  
}

// new channel
func NewChannel(conn *amqp.Connection) *amqp.Channel {  

	cha, err := conn.Channel() 
	if err != nil {
		FailOnError(err, "Failed to open a channel")
	}
	return cha
}

// new queue
func NewQueue(queue string, ch *amqp.Channel) amqp.Queue{  

	q, err := ch.QueueDeclare(
		queue, // name of the queue
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	  )
	
	if err != nil { 
		FailOnError(err, "Failed to declare a queue")  
	}
	return q
}

// publish to queue
func SendMessage(q_name string, body string){ 

	conn := Connect()
	cha := NewChannel(conn)
	q := NewQueue(q_name, cha)
 
	err := cha.Publish(
	  "",     // exchange
	  q.Name, // routing key
	  false,  // mandatory
	  false,  // immediate
	  amqp.Publishing {
		ContentType: "text/plain",
		Body:        []byte(body),
	  })

	if err != nil { 
		FailOnError(err, "Failed to publish a message")
	}


	fmt.Println("Sending ...", body)
	
	defer conn.Close() 
	defer cha.Close()
}
