package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"net/http"
	"os"
	"time"
)

func main() {

	//Establish RabbitMQ connection
	conn, err := amqp.Dial(os.Getenv("AMPQConn"))
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"eth",    // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")
	////////////////////////////////////////////////////////////////////

	for {

		lastBlockNumber := readLastBlock()
		nextBlockNumber := increamentHex(lastBlockNumber)
		//Create Infura API query
		params := setParams(nextBlockNumber, true)
		data := Payload{Jsonrpc: "2.0", Method: "eth_getBlockByNumber", Params: params, ID: 1}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
		}
		body := bytes.NewReader(payloadBytes)

		url := "https://mainnet.infura.io/v3/" + os.Getenv("INFURA_API_KEY")
		req, err := http.NewRequest("POST", url, body)
		if err != nil {
		}
		req.Header.Set("Content-Type", "application/json")
		//		fmt.Println(req)
		result := handleRequest(req)

		//Parse Response and send message over RabbitMQ
		block := processBlock(result)
		if block.Number == "" {
			fmt.Println(fmt.Sprintf("Last block seen %v", lastBlockNumber))
			time.Sleep(5 * time.Second)
			continue
		}
		writeLastBlock(block.Number)

		payments := processTxs(block.Transactions)
		for i := range payments {
			payment, err := json.Marshal(payments[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			err = ch.Publish(
				"eth",      // exchange
				"payments", // routing key
				false,      // mandatory
				false,      // immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(payment),
				})
			//			log.Printf(" [x] Sent %s", payments[i].String())
			failOnError(err, "Failed to publish a message")
		}
	}

}
