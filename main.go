package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("hello bitcoin")

	//block := NewBlock("The Times 2009-1-03 Chancellor on brink of second bailout for banks", []byte{0x0000000000000000})

	bc := NewBlockChain()

	bc.AddBlock("俺是第一个block")

	bc.AddBlock("俺是第二个block")

	bc.AddBlock("俺是第三个block")

	for idx, block := range bc.Blocks {
		fmt.Printf("======================== %d ===================\n", idx)
		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("PreBlockHash : %x\n", block.PreBlockHash)
		fmt.Printf("MerkleRoot : %x\n", block.MerkleRoot)

		timeFormat := time.Unix(int64(block.TimeStamp), 0).Format("2006-01-02 15:04:05")

		fmt.Printf("TimeStamp : %s\n", timeFormat)
		fmt.Printf("Difficulity : %d\n", block.Difficulity)
		fmt.Printf("Nonce : %d\n", block.Nonce)
		fmt.Printf("Version : %d\n", block.Version)
		fmt.Printf("hash : %x\n", block.Hash)
		fmt.Printf("data : %s\n", block.Data)

		pow := NewProofOfWork(block)
		fmt.Printf("Isvaild: %v\n", pow.IsVaild())
	}

}
