package messages 

import (
	"log"
	"strconv"
	"time"
 
	"github.com/vonmutinda/crafted/api/database"
)


// consume messages 
func Consume(){

	conn := Connect()  
	ch := NewChannel(conn)
	q := NewQueue("updated_at", ch)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		FailOnError(err, "Failed to establish consumer!")
	}
 
	forever := make(chan bool)

	go func() {
	  for d := range msgs {
		log.Printf("Received a message: %s", d.Body) 
		aid, err := strconv.ParseInt(string(d.Body), 10, 32)
		if err != nil {
			log.Println("Error parsing message body")
		} 

		
 
		database.GetDB().Exec(`
			UPDATE articles
			SET updated_at=?
			WHERE id=?
			`,
			time.Now(),
			aid,
		) 
		
  
	  }
	}() 

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
 
}
 
  
