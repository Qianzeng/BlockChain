package BLC

type Blockchain struct {
	Blocks []*Block
}

func CreateBlockchainWithGenesisblock() *Blockchain {
	genesisBlock := CreateGenesisBlock("Genesis Data.......")
	// 返回区块链对象
	return &Blockchain{[]*Block{genesisBlock}}
}