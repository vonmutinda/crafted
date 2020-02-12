## Crafted - All about Golang & Microservices
Let's built a ROBUST microservice as possible. Bringing on board all
possible awesome technologies on board.

I like simple and that's exactly how we'll do this. ```go #LetThereBeEXCELLENCE```


## Description
Just for func

```go
    go func(){

    }(yourself)
```

Inspired By [`Washington Redskins`](https://en.wikipedia.org/wiki/Go_Fund_Yourself)

## Features
- [x] GET,POST,PUT,DELETE users and articles
- [x] DB connection with ```gorm```
- [x] RAW SQL usage alongside [GORM](https://gorm.io)
- [x] CLI tooling with Cobra 

## #TODOs
- [ ] Unit Testing
- [ ] Authentication and Authorization Middlewares
- [ ] GraphQL APIs 
- [ ] gRPC 
- [ ] Docker (containerization)
- [ ] Service Deployment

## URLS

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

- If you are using a different `db` from `postgres`, make sure you import its corresponding dialect in `package database`
- Create `db` and it's name in .env file.
- Run `go run main.go crafted` or `go build && ./crafted crafted`

## Technologies Used 
Here's a list of technologies used in this project

- [Golang version `go1.13.7`](https://golang.org)
- [Cobra](https://github.com/spf13/cobra)
- [gorilla/mux HTTP framework](https://github.com/gorilla/mux). You could as well use [Gin](https://github.com/gin-gonic/gin)
- [Gorm ORM](https://gorm.io/). However I'd highly recommend writing raw SQL.

## NOTES:
- [x] Use  `SentenceCase` when naming funcs intended for global usage.
- [x] Receivers must be pointers.
- [x] Channels are used when Feedback is expected from a `go routine`
- [x] Waitgroups are used when we don't care about Feedback. We only want the job done. 
- [x] A WaitGroup is of type `sync.WaitGroup` but we use a pointer of that type in a `go routine`. 

## Resources 
Below are helpful resources on where to read more about `Go/Golang`.
- [Go docs](https://tour.golang.org/)
- [Go by Example](https://gobyexample.com/)
- [Nic Jackson](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)'s Tutorial Series
  



  <!-- http://tumregels.github.io/Network-Programming-with-Go/ -->