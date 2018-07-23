package main

import "fmt"
import "./BLC"

func main() {
	//block := BLC.NewBlock("Genenis is Block", 1, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	genesisBlock := BLC.CreateGenesisBlock("Genesis Block")
	fmt.Println(genesisBlock)
	fmt.Println("Qzeng")
}
