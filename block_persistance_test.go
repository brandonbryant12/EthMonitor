package main 

import(
  "testing"
  "fmt"
)

func Test_readBlock(t *testing.T){
	//Testing when there is no latestBlock.txt in memory
	expected := "0x5E9D59"
	actual := readLastBlock()
	
	if expected != actual{
	  t.Error(fmt.Sprintf("expected %v, actual %v", expected, actual))
	} 
}
/*	
func Test_writeLastBlock_and_readLastBlock(t *testing.T){
  expected := "0xB34BB6AA"
  writeLastBlock(expected)
  actual := readLastBlock()
   if expected != actual{
        t.Error(fmt.Sprintf("expected %v, actual %v", expected, actual))
    }

}
*/
func Test_writeLastBlock_and_readLastBlock_with_lowercase(t *testing.T){
  expected := "0x5E9D59"
  writeLastBlock(expected)
  actual := readLastBlock()
   if expected != actual{
        t.Error(fmt.Sprintf("expected %v, actual %v", expected, actual))
    }

}

	
func Test_increamentHex(t *testing.T){
  original := "0xbbb6aa"
  expected := "0xbbb6ab"
  actual := increamentHex(original)
  if expected != actual{
    t.Error(fmt.Sprintf("expected %v, actual %v", expected, actual))
  }   
}
