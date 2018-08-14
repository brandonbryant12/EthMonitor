package main

import "testing"
func TestProcessTxs(t *testing.T) {
	var txs = []Transaction{
			Transaction{
				BlockHash: "0x174c96af107758036879762866ad1854bc35024be58c86578e3772735eb1e683",
				BlockNumber: "0x5d7049",
       				Gas:"0x15f90",      
       				GasPrice:"0x6fc23ac00" ,        
        			Hash:"0x251ced8e0ec34e4fd95381655c20fca3f9c2c7f3de7e8cb4c3f7de71b545ba0e",         
       	 			Input:"0xa9059cbb000000000000000000000000876eabf441b2ee5b5b0554fd502a8e0600950cfa00000000000000000000000000000000000000000000002cc70d7faa18a80000",           
        			Nonce:"0x1",           
        			To:"0x7025bab2ec90410de37f488d1298204cd4d6b29d",           
        			R: "0xd8e7d50baca8853f75a520e7b8d5e0280bfb76539237f361f0cff38dda236790",       	        
        			S:"0x463f43fc92ca70cf5a222bdcbb427acfd9f5708e4ba7b54cc4cae245256629d5" ,                
        			TransactionIndex:"0x11", 
        			V:"0x25",                
       	 			Value:"0x0",
			 },
		}
	var payments = processTxs(txs)
	if payments[0].Hash != "0x251ced8e0ec34e4fd95381655c20fca3f9c2c7f3de7e8cb4c3f7de71b545ba0e" {
		t.Error("expected 0x251ced8e0ec34e4fd95381655c20fca3f9c2c7f3de7e8cb4c3f7de71b545ba0e got, ",  payments[0].Hash)
		
	}
	
}

