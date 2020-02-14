## Crafted - All about Golang & Microservices
Let's built a microservice as ROBUST as possible. Bringing on-board all
possible awesome technologies.

I like simple and that's exactly how we'll do this. ```#LetThereBeEXCELLENCE```. 
Collaboration is totally welcome. It's a great way of learning new stuff along the way.


## Description
Just for func

```go
    go func(){

    }(yourself)
```

Inspired By [`Washington Redskins`](https://en.wikipedia.org/wiki/Go_Fund_Yourself)

## Features
- [x] RESTful operations
- [x] DB connection with ```gorm``` and RAW SQL usage alongside [GORM](https://gorm.io)
- [x] CLI tooling with Cobra 
- [x] Queues and Messaging with `RabbitMQ`
- [x] Logging Errors

## #TODOs
- [ ] Unit Testing
- [ ] Authentication and Authorization Middlewares
- [ ] GraphQL APIs 
- [ ] gRPC 
- [ ] Docker (containerization) 
- [ ] Service Deployment
- [ ] Grafana and Prometheus Integrations (Later on)

## Setup Local
For set up on your machine .
- Clone the repo `git clone https://github.com/vonmutinda/crafted.git`.
- Run `go mod init` to check if go modules is already initialized.
- Touch  `.env` file and paste the following configurations.
  
### .env file

```go
    PORT=":9000"
    DB_DRIVER="postgres" # <-provide your own-->
    DB_HOST="localhost"
    DB_PORT="5432"
    DB_USER="username" # <-provide your own-->
    DB_NAME="db_name" # <-provide your own-->
    DB_PASS="db_pass" # <-provide your own-->
```

If you are using a different `db` from `postgres`, make sure you import its corresponding dialect in `package database`

- MySQL
```go
package database

import ( 
	"github.com/jinzhu/gorm" 
	_"github.com/jinzhu/gorm/dialects/mysql"
)
// rest of code goes here
```

Create `db` and add it's name in ```.env``` file.
Run `go run main.go crafted` or `go build && ./crafted crafted`

Install [Docker](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-18-04)

Start RabbitMQ container 
```cmd 
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
```

Since we constantly want to listen for any messages hitting the queue, run the `cobra command`
```cmd 
go run main.go consume
```

- Logging
You realise this application is growing too big and once users begin interacting with it in production,
We'll need a way to know where it fails. 

Later on we'll configure our logger and other parts of our app with [`Prometheus`](https://prometheus.io/) are [Grafana](https://grafana.com/) which are more like [Sentry](https://sentry.io/welcome/) and [Google Analytics](https://analytics.google.com/analytics/web/)


## Technologies Used 
Here's a list of technologies used in this project

- [Golang version `go1.13.7`](https://golang.org)
- [Cobra](https://github.com/spf13/cobra)
- [gorilla/mux HTTP framework](https://github.com/gorilla/mux). You could as well use [Gin](https://github.com/gin-gonic/gin)
- [Gorm ORM](https://gorm.io/). However I'd highly recommend writing raw SQL.
- [RabbitMQ](https://www.rabbitmq.com/tutorials/tutorial-one-go.html) Messaging and Queues. 

## NOTES:
- [x] Use  `Sentence case` when naming funcs intended for global usage.
- [x] Receivers must be pointers.
- [x] Channels are used when Feedback is expected from a `go routine`
- [x] Waitgroups are used when we don't care about Feedback. We only want the job done. 
- [x] A WaitGroup is of type `sync.WaitGroup` but we use a pointer of that type in a `go routine`. 

## Resources 
Below are helpful resources on where to read more about `Go/Golang`.
- [Go docs](https://tour.golang.org/)
- [Go by Example](https://gobyexample.com/)
- [Nic Jackson](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)'s Tutorial Series
