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

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("BlockBucket"))
		if b != nil {
			data := b.Get([]byte("l"))
			fmt.Printf("%s\n",data)
			data = b.Get([]byte("ll"))
			fmt.Printf("%s\n",data)
		}
		return nil
	})

	if err !=nil {
		log.Panic(err)
	}
}