package main

import (
	"fmt"
	"math/big"
)

type Payment struct {
	Currency string
	Address  string
	Amount   *big.Float
	Hash     string
}

func (payment *Payment) String() string {
	return fmt.Sprintf("Currency: %v\nAddress: %v\nAmount: %v\nHash: %v", payment.Currency, payment.Address, payment.Amount, payment.Hash)
}
