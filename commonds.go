package main

import (
	"bytes"
	"fmt"
	"time"
)

func (cli *CLI) AddBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Printf("添加区块成功！！\n")
}

func (cli *CLI) PrintChain() {
	it := cli.bc.NewIterator()

	for {
		block := it.Next()
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

		if bytes.Equal(block.PreBlockHash, []byte{}) {
			fmt.Printf("区块链遍历结束～～~!\n")
			break
		}

	}
}
