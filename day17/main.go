package main

import (
	"./BLC"
)

func main(){
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()

	blockchain.AddBlockToBlockchain("Send 100RMB To qzeng")
	blockchain.AddBlockToBlockchain("Send 1000RMB To QZ")
	blockchain.AddBlockToBlockchain("Send 5000RMB To qianzeng")
	blockchain.AddBlockToBlockchain("Send 100RMB To Q")

	blockchain.Printchain()

}
