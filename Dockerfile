FROM golang:alpine
 
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app 

COPY . . 

RUN go mod download


RUN go build -o main .

EXPOSE 9000 

CMD [  "./main", "crafted" ]