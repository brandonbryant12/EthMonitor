FROM golang:1.8
RUN mkdir -p /go/src/app
WORKDIR /go/src/
COPY . .

RUN go get github.com/streadway/amqp
RUN go build main.go read_write_block.go Payload.go Payment.go Block.go Transaction.go helpers.go receive.go
