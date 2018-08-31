package BLC

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
	"math/big"
	"time"
)

const dbName ="blockchain.db"
const blockTableName = "blocks"

type Blockchain struct {
	Tip []byte
	DB *bolt.DB
}

func (blockchain *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{blockchain.Tip,blockchain.DB}
}
func (blc *Blockchain) Printchain() {
	//var block *Block
	//var currentHash []byte = blc.Tip
	blockchainIterator := blc.Iterator()
	for {
		block := blockchainIterator.Next()
		fmt.Printf("Height：%d\n",block.Height)
		fmt.Printf("PrevBlockHash：%x\n",block.PrevBlockHash)
		fmt.Printf("Data：%s\n",block.Data)
		fmt.Printf("Timestamp：%s\n",time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash：%x\n",block.Hash)
		fmt.Printf("Nonce：%d\n",block.Nonce)


		fmt.Println()

		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0{
			break;
		}

	}
}

// 增加区块到区块链里面
func (blc *Blockchain) AddBlockToBlockchain(data string)  {
	// 创建新区块
	err := blc.DB.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockTableName))
		if b!=nil {
			blockBytes := b.Get(blc.Tip)
			block := DeserializeBlock(blockBytes)
			newBlock := NewBlock(data,block.Height+1,block.Hash)
			err := b.Put(newBlock.Hash,newBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}
			err = b.Put([]byte("l"),newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			blc.Tip = newBlock.Hash
		}
		return nil
	})

	if err != nil {
		log.Panic(err)
	}
}


//1. 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	// 创建或者打开数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var blockHash []byte

	err = db.Update(func(tx *bolt.Tx) error{

		//  获取表
		b := tx.Bucket([]byte(blockTableName))

		if b == nil {
			// 创建数据库表
			b,err = tx.CreateBucket([]byte(blockTableName))

			if err != nil {
				log.Panic(err)
			}
		}



		if b != nil {
			// 创建创世区块
			genesisBlock := CreateGenesisBlock("Genesis Data.......")
			// 将创世区块存储到表中
			err := b.Put(genesisBlock.Hash,genesisBlock.Serialize())
			if err != nil {
				log.Panic(err)
			}

			// 存储最新的区块的hash
			err = b.Put([]byte("l"),genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			blockHash = genesisBlock.Hash
		}

		return nil
	})



	// 返回区块链对象
	return &Blockchain{blockHash,db}
}