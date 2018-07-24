package BLC

import (
	"time"
)
type Block struct {
	Height        int64
	PrevBlockHash []byte
	Data          []byte
	Timestamp     int64
	Hash          []byte
	Nonce         int64
}

func NewBlock(data string,height int64,prevBlockHash []byte) *Block {

	//创建区块
	block := &Block{height,prevBlockHash,[]byte(data),time.Now().Unix(),nil,0}
	// 工作量证明
	pow := NewProofOfWork(block)

	hash,nonce := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block

}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}

