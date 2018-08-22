package main

import (
    "io/ioutil"
    "strconv"
    "fmt"
"strings"
)


func writeLastBlock(s string){
  d1 := []byte(s)
  err := ioutil.WriteFile("lastBlock.txt", d1, 0644)
  check(err)
}

func readLastBlock() string{
  
  if Exists("lastBlock.txt"){
    dat, err := ioutil.ReadFile("lastBlock.txt")
    check(err)
    return strings.TrimSpace(string(dat))

  }
    return "latest"
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
  n2 = n2+1
  return fmt.Sprintf("0x%v",strconv.FormatInt(n2, 16))
}
