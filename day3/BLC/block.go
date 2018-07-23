package BLC

import (
	"fmt"
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
)
type Block struct {
	Height int64
	PrevBlockHash []byte
	Data []byte
	Timestamp int64
	Hash []byte
}

func (block *Block) SetHash() {
	heightBytes := IntToHex(block.Height)
	fmt.Println(heightBytes)
	//2进制str
	timeString := strconv.FormatInt(block.Timestamp, 2)
	fmt.Println(timeString)
	//2进制字节数组
	timeBytes := []byte(timeString)
	fmt.Println(timeBytes)
	//拼接属性
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		block.PrevBlockHash,
		block.Data,
		timeBytes}, 
		[]byte{})

	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
}

func NewBlock(data string,height int64,prevBlockHash []byte) *Block {

	//创建区块
	block := &Block{height,prevBlockHash,[]byte(data),time.Now().Unix(),nil}

	//设置Hash
	block.SetHash()

	return block

}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(data,1,[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}

