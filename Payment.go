package main

import (
	"fmt"
	"math/big"
)

type Payment struct {
	Currency string
	Address  string
	Amount   big.Float
	Hash     string
}

func (payment *Payment) String() string {
	return fmt.Sprintf("currency: %v\naddress: %v\namount: %v\nhash: %v", payment.Currency, payment.Address, payment.Amount.String(), payment.Hash)
}
