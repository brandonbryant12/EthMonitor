package main

import (
	//"math/big"
	"testing"
)

func TestProcessTxs(t *testing.T) {
	var payments = processTxs(txs)
	if payments[0].Hash != "0x251ced8e0ec34e4fd95381655c20fca3f9c2c7f3de7e8cb4c3f7de71b545ba0e" {
		t.Error("expected 0x251ced8e0ec34e4fd95381655c20fca3f9c2c7f3de7e8cb4c3f7de71b545ba0e got, ", payments[0].Hash)
	}
	if payments[0].Amount.String() != "0" {
		t.Error("expected 0, got: ", payments[0].Amount.String())
	}
	if payments[0].Address != "0x7025bab2ec90410de37f488d1298204cd4d6b29d" {
		t.Error("expected 0x7025bab2ec90410de37f488d1298204cd4d6b29d, got: ", payments[0].Address)
	}
	if payments[0].Currency != "ETH" {
		t.Error("expected ETH, got: ", payments[0].Currency)
	}
}
