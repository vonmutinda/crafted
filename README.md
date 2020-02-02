# Crafted - All about Golang & Microservices

## Description
- Just for func

## URLS

## Setup Local
- For set up on your machine .
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
- Run `go run main.go` or `go build && ./crafted`

## NOTES:
- [x] Use  `SentenceCase` when naming funcs intended for global usage.
- [x] 