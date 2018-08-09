package BLC

import (
	"math/big"
	"crypto/sha256"
	"fmt"
	"bytes"
)

const targetBit = 16 //256 位HASH前必须有16个0  难度

type ProoFOfWork struct {
	Block *Block
	target *big.Int // 大数据存储  工作难度
}

// 0000 0000 0000 0000 1011 0002 .... 0001  （256位） 通过targetBit
func (pow *ProoFOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{}, 
	)

	return data
}
// 判断区块有效性。
func (prooFOfWork * ProoFOfWork) IsValid() bool {
	var hashInt big.Int
	hashInt.SetBytes(prooFOfWork.Block.Hash)

	if prooFOfWork.target.Cmp(&hashInt) == 1 {
		return true
	}
	return false
}
func (prooFOfWork * ProoFOfWork) Run() ([]byte, int64) {

	//1. 将Block的属性拼接成字节数组

	//2. 生成hash

	//3. 判断hash有效性，如果满足条件，跳出循环

	nonce := 0

	var hashInt big.Int // 存储我们新生成的hash
	var hash [32]byte

	for {
		//准备数据
		dataBytes := prooFOfWork.prepareData(nonce)

		// 生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x",hash)


		// 将hash存储到hashInt
		hashInt.SetBytes(hash[:])
		//判断hashInt是否小于Block里面的target
		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		if prooFOfWork.target.Cmp(&hashInt) == 1 {
			break
		}

		nonce = nonce + 1
	}

	return hash[:],int64(nonce)
}
// 假如 难度我要求256 前边只是有2个0    那么我需要让target值 左移 256-2 位 这样 0100 ....0000(256位)
// 我们假如哈希8位 还是前边有两个0 8-2 =6 左移 6位 相当于 2的6次幂 值位64
// 如何证明我们得出的HASH 满足条件 很简单  只要与 0100 0000   比较前两位 0010 0000 这种结构即可 换句话说 只要小于 0100 0000的HASN就可以。
//256 位 需要前边16个0 那么 我们计算 2的（256-16）次幂的值，小于这个就可以
func NewProofOfWork(block  *Block) *ProoFOfWork{
	//pow 算法 1 初始化 target 1   0000  .... 0001
	target := big.NewInt(1)

	 //pow 算法 2  左移256 - targetBit
	target = target.Lsh(target,256 -targetBit)

	return &ProoFOfWork{block, target}
}