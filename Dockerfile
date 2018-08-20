FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get github.com/streadway/amqp
RUN go build ethAPI.go Payload.go Payment.go Block.go Transaction.go helpers.go
RUN go build recieve.go
