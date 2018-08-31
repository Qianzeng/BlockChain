package main

import (
	"github.com/boltdb/bolt"
	"log"
	"fmt"
)

func main(){
	db,err :=bolt.Open("my.db",0600,nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b,err := tx.CreateBucket([]byte("BlockBucket"))
		if err != nil {
			return fmt.Errorf("create bucket:%s",err)
		}
		if b != nil {
			err := b.Put([]byte("l"),[]byte("Send 100 BTC To QZ....."))
			if err !=  nil {
				log.Panic("shujucunshibai")
			}
		}
		return nil
	})

	if err !=nil {
		log.Panic(err)
	}
}