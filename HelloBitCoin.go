package main

import (
	"fmt"
)

type Block struct {
	PreBlockHash []byte //当前快哈希
	Hash         []byte //当前区块哈希
	Data         []byte //数据体
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		PreBlockHash: preBlockHash,
		Hash:         []byte{},
		Data:         []byte(data),
	}
	return &block
}

func main() {
	fmt.Println("hello bitcoin")

	block := NewBlock("啊啊啊", []byte{0x0000000000000000})
	fmt.Printf("prv : %x\n", block.PreBlockHash)
	fmt.Printf("hash : %x\n", block.Hash)
	fmt.Printf("data : %s\n", block.Data)
}
