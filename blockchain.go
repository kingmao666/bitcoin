package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

//从bolt中加载区块
type BlockChain struct {
	db   *bolt.DB
	tail []byte //最后一个区块的哈希值
}

const blockChainName = "blackChain.db"
const blockBucketName = "blockBucket"
const lastHashKey = "lastHashKey"

//创建区块链
func NewBlockChain() *BlockChain {
	//gen := NewBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks", []byte{0x0000000000000000})
	//bc := BlockChain{Blocks: []*Block{gen}}
	//return &bc

	db, err := bolt.Open(blockChainName, 0600, nil)

	if err != nil {
		log.Panic(err)
	}

	//defer db.Close()

	var tail []byte

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucketName))
		if b == nil {
			fmt.Printf("bucket不存在，准备创建！\n")
			b, err = tx.CreateBucket([]byte(blockBucketName))

			if err != nil {
				log.Panic(err)
			}
			gen := NewBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks", []byte{})

			b.Put([]byte(gen.Hash), gen.Serialize())
			b.Put([]byte(lastHashKey), gen.Hash)
			tail = gen.Hash
		} else {
			tail = b.Get([]byte(lastHashKey))
		}
		return nil
	})

	return &BlockChain{db, tail}

}

//序列化
func (block *Block) Serialize() []byte {
	//将block数据转换成字节流
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(block)

	if err != nil {
		fmt.Println("encode failed", err)
		os.Exit(1)
	}

	return buffer.Bytes()
}

//反序列化
func Deserialize(data []byte) *Block {
	//fmt.Printf("解码传入的数据： %x\n", data)
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}
	return &block
}

//添加区块
func (bc *BlockChain) AddBlock(data string) {
	bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucketName))
		if b == nil {
			fmt.Printf("bucket不存在，请检查数据！\n")
			os.Exit(1)
		}
		block := NewBlock(data, bc.tail)

		b.Put([]byte(block.Hash), block.Serialize())
		b.Put([]byte(lastHashKey), block.Hash)
		bc.tail = block.Hash
		return nil
	})
}

//定义一个区块链迭代器
type BlockChainIterator struct {
	db      *bolt.DB
	current []byte
}

//创建迭代器
func (bc *BlockChain) NewIterator() *BlockChainIterator {
	return &BlockChainIterator{bc.db, bc.tail}
}

func (it *BlockChainIterator) Next() *Block {
	var block Block

	it.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucketName))
		if b == nil {
			fmt.Printf("bucket不存在，请检查数据！\n")
			os.Exit(1)
		}
		blockInfo := b.Get(it.current)
		block = *Deserialize(blockInfo)
		it.current = block.PreBlockHash

		return nil
	})

	return &block

}
