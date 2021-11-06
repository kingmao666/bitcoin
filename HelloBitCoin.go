package main

import (
	"crypto/sha256"
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
	block.SetHash()
	return &block
}

func (block *Block) SetHash() {
	var data []byte
	data = append(data, block.PreBlockHash...)
	data = append(data, block.Data...)
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

//创建区块链，使用Block数组模拟
type BlockChain struct {
	Blocks []*Block
}

//创建区块链
func NewBlockChain() *BlockChain {
	gen := NewBlock("aaa", []byte{0x0000000000000000})
	bc := BlockChain{Blocks: []*Block{gen}}
	return &bc
}

//添加区块
func (bc *BlockChain) AddBlock(data string) {
	//1.拿到上一个区块
	lastBlock := bc.Blocks[len(bc.Blocks)-1]
	//2.创建区块
	newBlock := NewBlock(data, lastBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func main() {
	//fmt.Println("hello bitcoin")

	//block := NewBlock("The Times 2009-1-03 Chancellor on brink of second bailout for banks", []byte{0x0000000000000000})

	bc := NewBlockChain()

	bc.AddBlock("第一个新增block")

	for _, blockt := range bc.Blocks {
		fmt.Printf("prv : %x\n", blockt.PreBlockHash)
		fmt.Printf("hash : %x\n", blockt.Hash)
		fmt.Printf("data : %s\n", blockt.Data)
	}

}
