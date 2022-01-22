package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		b1 := tx.Bucket([]byte("bucket1"))
		if b1 == nil {
			b1, err = tx.CreateBucket([]byte("bucket1"))
			if err != nil {
				log.Panic(err)
			}
		}
		//bucket创建完毕
		//写数据使用put，读数据使用get
		err = b1.Put([]byte("name1"), []byte("bojackworkman"))
		if err != nil {
			log.Printf("写入数据失败 name1：bojackworkman！\n")
		}

		err = b1.Put([]byte("name2"), []byte("kingmao666"))
		if err != nil {
			log.Printf("写入数据失败 name2：kingmao666！\n")
		}

		name1 := b1.Get([]byte("name1"))
		name2 := b1.Get([]byte("name2"))
		name3 := b1.Get([]byte("name3"))

		fmt.Printf("name1: %s\n", name1)
		fmt.Printf("name2: %s\n", name2)
		fmt.Printf("name3: %s\n", name3)

		return nil
	})

}
