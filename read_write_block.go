package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func writeLastBlock(s string) {
	d1 := []byte(s)
	err := ioutil.WriteFile("/data/lastBlock.txt", d1, 0644)
	check(err)
}

func readLastBlock() string {

	if Exists("/data/lastBlock.txt") {
		dat, err := ioutil.ReadFile("/data/lastBlock.txt")
		check(err)
		if string(dat) == "" {
			return "0x5F1184"
		}
		return strings.TrimSpace(string(dat))

	}
	fmt.Println("Does not exist")
	var file, err = os.Create("/data/lastBlock.txt")
	check(err)
	defer file.Close()
	return "0x5F1184"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func increamentHex(s string) string {
	n, err := strconv.ParseUint(s[2:], 16, 32)
	if err != nil {
		panic(err)
	}
	n2 := int64(n)
	n2 = n2 + 1
	return fmt.Sprintf("0x%v", strconv.FormatInt(n2, 16))
}
