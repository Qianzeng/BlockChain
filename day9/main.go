package main

import "fmt"
import "./BLC"
func main() {
	//blockchain := BLC.CreateBlockchainWithGenesisblock()
	
	//blockchain.AddBlockToBlockchain("Send 100RMB To zhangqiang",blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1,blockchain.Blocks[len(blockchain.Blocks) - 1].Hash)
	//blockchain.AddBlockToBlockchain("Send 200RMB To qzeng",blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1,blockchain.Blocks[len(blockchain.Blocks) - 1].Hash)
	//blockchain.AddBlockToBlockchain("Send 300RMB To QZ",blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1,blockchain.Blocks[len(blockchain.Blocks) - 1].Hash)
	//blockchain.AddBlockToBlockchain("Send 500RMB To ZQ",blockchain.Blocks[len(blockchain.Blocks) - 1].Height + 1,blockchain.Blocks[len(blockchain.Blocks) - 1].Hash)

	//fmt.Println(blockchain)
	//fmt.Println(blockchain.Blocks)
	block := BLC.NewBlock("Test",1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%x\n",block.Hash)
	bytes:= block.Serialize()

	block = BLC.DeserializeBlock(bytes)
	fmt.Printf("%d\n",block.Nonce)
	fmt.Printf("%x\n",block.Hash)
	fmt.Println("Qzeng\n\r")
}
