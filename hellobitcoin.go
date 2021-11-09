package main

import (
	"fmt"
)

func main() {
	//fmt.Println("hello bitcoin")

	//block := NewBlock("The Times 2009-1-03 Chancellor on brink of second bailout for banks", []byte{0x0000000000000000})

	bc := NewBlockChain()

	bc.AddBlock("俺是第一个block")

	bc.AddBlock("俺是第二个block")

	bc.AddBlock("俺是第三个block")

	for idx, blockt := range bc.Blocks {
		fmt.Printf("======================== %d ===================\n", idx)
		fmt.Printf("prv : %x\n", blockt.PreBlockHash)
		fmt.Printf("hash : %x\n", blockt.Hash)
		fmt.Printf("data : %s\n", blockt.Data)
	}

}
