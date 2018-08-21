package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"
)

func hexToEth(s string) big.Float {
	if s == "0x0" {
		return *big.NewFloat(0)
	}
	weiInt := new(big.Int)
	weiInt, ok := weiInt.SetString(s[2:], 16)
	if !ok {
		fmt.Println("SetString: error")
	}
	wei := new(big.Float).SetInt(weiInt)
	ratio := big.NewFloat(.000000000000000001)
	eth := new(big.Float).Mul(wei, ratio)
	return *eth
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

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
