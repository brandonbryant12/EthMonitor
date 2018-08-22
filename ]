package main 

import(
  "testing"
  "fmt"
)

func Test_writeLastBlock_and_readLastBlock(t *testing){
  lastBlockInHex := ""
  expected := "0xBBB6AA"
  writeLastBlock(expected)
  actual = readLastBlock()
   if expected != actual{
        t.Error(fmt.Sprintf("expected %v, actual %v", expected, actual))
    }

}

func Test_writeLastBlock_and_readLastBlock_with_lowercase(t *testing){
  lastBlockInHex := ""  
  expected := "0xbbb6aa"
  writeLastBlock(expected)
  actual = readLastBlock()
   if expected != actual{
        t.Error(fmt.Sprintf("expected %v, actual %v", expected, actual))
    }

}



func Test_getNextBlockInit(t *testing.T){
	//Testing when there is no latestBlock.txt in memory
	expected := "latest"
	actual := getNextBlock()
	
	if expected != actual{
	  t.Error(fmt.Sprintf("expected %v, actual %v", expected, actual))
	} 

		
}
	

