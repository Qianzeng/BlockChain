package BLC

type Blockchain struct {
	Blocks []*Block
}

func (blc *Blockchain) AddBlockToBlockchain(data string, height int64, preHash []byte) {

	newBlock := NewBlock(data, height, preHash)

	blc.Blocks = append(blc.Blocks, newBlock)
}

func CreateBlockchainWithGenesisblock() *Blockchain {
	genesisBlock := CreateGenesisBlock("Genesis Data.......")
	// 返回区块链对象
	return &Blockchain{[]*Block{genesisBlock}}
}