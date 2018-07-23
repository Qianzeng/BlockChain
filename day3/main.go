package main

import "fmt"
import "./BLC"

func main() {
	genesisBlockchain := BLC.CreateBlockchainWithGenesisblock()
	fmt.Println(genesisBlockchain)
	fmt.Println(genesisBlockchain.Blocks)
	fmt.Println(genesisBlockchain.Blocks[0])
	fmt.Println("Qzeng")
}
