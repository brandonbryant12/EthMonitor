FROM golang:1.8

WORKDIR /usr/local/go/src/app
COPY . .

RUN go get github.com/streadway/amqp
RUN go build main.go Payload.go Payment.go Block.go Transaction.go helpers.go
#RUN go build receive.go
