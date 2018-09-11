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

var coins = map[string]string{
"0xB8c77482e45F1F44dE1745F52C74426C631bDD52":"BNB",
"0xd850942ef8811f2a866692a623011bde52a462c1":"VEN",
"0xd26114cd6EE289AccF82350c8d8487fedB8A0C07":"OMG",
"0xe41d2489571d322189246dafa5ebde1f4699f498":"ZRX",
"0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2":"MKR",
"0xb5a5f22694352c15b00323844ad545abb2b11028":"ICX",
"0x5ca9a71b1d01849c0a95490cc00559717fcf0d1d":"AE",
"0xa15c7ebe1f07caf6bff097d8a589fb8ac49ae5b3":"NPXS",
"0x1985365e9f78359a9B6AD760e32412f4a445E862":"AUG",
"0x0d8775f648430679a709e98d2b0cb6250d2887ef":"BAT",
"0xa74476443119A942dE498590Fe1f2454d7D4aC0d":"GNT",
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
		if(len(txs[i].Input) == 138 ){
				//If token transfer
			if(txs[i].Input[2:34] == "a9059cbb000000000000000000000000"){
				payments = append(payments, Payment{
					Currency: fmt.Sprintf("%v.%v",coins[txs[i].To],txs[i].To),
					Address: fmt.Sprintf("0x%v",txs[i].Input[34:74]),
					Amount: hexToEth(fmt.Sprintf("0x%v",txs[i].Input[75:len(txs[i].Input)])),
					Hash: txs[i].Hash})
			}
			continue
		}
		payments = append(payments, Payment{
			Currency: "ETH",
			Address:  txs[i].To,
			Amount:   hexToEth(txs[i].Value),
			Hash:     txs[i].Hash})
			
	}

	return payments
}
