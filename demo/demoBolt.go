package main

import (
	"log"

	"github.com/boltdb/bolt"
)

func main() {

	log.Println("sdfsdf")

	db, err := bolt.Open("test.db", 0600, nil)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

}
