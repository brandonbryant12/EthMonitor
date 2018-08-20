package main

import "fmt"

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
