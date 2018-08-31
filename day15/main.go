package main

import (
	"./BLC"
)

func main(){
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()

	}

