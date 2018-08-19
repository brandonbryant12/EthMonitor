package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"
)

func trimLeftChars(s string, n int) string {
	m := 0
	for i := range s {
		if m >= n {
			return s[i:]
		}
		m++
	}
	return s[:0]
}

func hexToEth(s string) *big.Float {
	if s == "0x0" {
		return big.NewFloat(0)
	}
	weiInt := new(big.Int)
	weiInt, ok := weiInt.SetString(s[2:], 16)
	if !ok {
		fmt.Println("SetString: error")
	}
	wei := new(big.Float).SetInt(weiInt)
	ratio := big.NewFloat(.000000000000000001)
	eth := new(big.Float).Mul(wei, ratio)
	return eth
}

type Payment struct {
	Currency string
	Address  string
	Amount   *big.Float
	Hash     string
}

type Payload struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

type Block struct {
	Difficulty       string        `json:"difficulty"`
	ExtraData        string        `json:"extraData"`
	GasLimit         string        `json:"gasLimit"`
	GasUsed          string        `json:"gasUsed"`
	Hash             string        `json:"hash"`
	LogsBloom        string        `json:"logsBloom"`
	Miner            string        `json:"miner"`
	MixHash          string        `json:"mixHash"`
	Nonce            string        `json:"nonce"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     string        `json:"receiptsRoot"`
	SHA3Uncles       string        `json:"sha3Uncles"`
	Size             string        `json:"size"`
	StateRoot        string        `json:"stateRoot"`
	Timestamp        string        `json:"timeStamp"`
	TotalDifficulty  string        `json:"totalDifficulty"`
	Transactions     []Transaction `json:"transactions"`
	TransactionsRoot string        `json:"transactions_root"`
	Uncles           []string      `json:"uncles"`
}

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	R                string `json:"r"`
	S                string `json:"s"`
	TransactionIndex string `json:"transactionIndex"`
	V                string `json:"v"`
	Value            string `json:"value"`
}

func (tx *Transaction) String() string {
	return fmt.Sprintf("to: %v\ngas: %v\ngasPrice: %v\nvalue: %v\nblockHash: %v\nblockNumber: %v\nhash: %v\ninput: %v\nnounce: %v\nr: %v\ns: %v\nv: %v\ntransactionIndex: %v\n",
		tx.To,
		tx.Gas,
		tx.GasPrice,
		tx.Value,
		tx.BlockHash,
		tx.BlockNumber,
		tx.Hash,
		tx.Input,
		tx.Nonce,
		tx.R,
		tx.S,
		tx.V,
		tx.TransactionIndex,
	)
}

func (payment *Payment) String() string {
	return fmt.Sprintf("Currency: %v\nAddress: %v\nAmount: %v\nHash: %v", payment.Currency, payment.Address, payment.Amount, payment.Hash)
}

//Refactor to make this method take an array of any type and size.
func setParams(blockNumber string, verbose bool) []interface{} {
	params := make([]interface{}, 2)
	params[0] = blockNumber
	params[1] = verbose
	return params
}

func handleRequest(req *http.Request) string {

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)

	s := strings.SplitAfter(string(content), `"result":`)[1]
	return s[:len(s)-len("}")]
}

//Takes and array of Transaction objects and outputs an array of Payment structs
func processTxs(txs []Transaction) []Payment {
	var payments []Payment
	for i := range txs {
		payments = append(payments, Payment{
			Currency: "ETH",
			Address:  txs[i].To,
			Amount:   hexToEth(txs[i].Value),
			Hash:     txs[i].Hash})
	}
	return payments
}

func processBlock(rawResponse string) Block {
	var block Block
	json.Unmarshal([]byte(rawResponse), &block)
	return block
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

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
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	//Create Infura API query
	params := setParams("latest", true)
	data := Payload{Jsonrpc: "2.0", Method: "eth_getBlockByNumber", Params: params, ID: 1}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		// handle err
	}

	var latestBlockHash = ""
	for {
		body := bytes.NewReader(payloadBytes)

		url := "https://mainnet.infura.io/v3/" + os.Getenv("INFURA_API_KEY")
		req, err := http.NewRequest("POST", url, body)
		if err != nil {
			// handle err
		}
		req.Header.Set("Content-Type", "application/json")
		result := handleRequest(req)

		//Parse Response and send message over RabbitMQ
		block := processBlock(result)
		if block.Hash == latestBlockHash {
			fmt.Println("Duplicate block")
			time.Sleep(5 * time.Second)
			continue
		}
		latestBlockHash = block.Hash

		payments := processTxs(block.Transactions)
		for i := range payments {
			payment, err := json.Marshal(payments[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			err = ch.Publish(
				"eth", // exchange
				"",    // routing key
				false, // mandatory
				false, // immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(payment),
				})
			//	log.Printf(" [x] Sent %s", payments[i].String())
			failOnError(err, "Failed to publish a message")
		}

		time.Sleep(5 * time.Second)
	}

}
