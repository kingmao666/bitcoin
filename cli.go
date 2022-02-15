package main

import (
	"fmt"
	"os"
)

const Usage = `
	./blockChain addBlock "xxxx" 添加数据到区块链
	./blockChain printChain      打印区块链
`

type CLI struct {
	bc *BlockChain
}

func (cli *CLI) Run() {
	cmds := os.Args
	if len(cmds) < 2 {
		fmt.Printf(Usage)
		os.Exit(1)
	}
	switch cmds[1] {
	case "addBlock":
		fmt.Printf("添加区块命令被调用，数据：%s\n", cmds[2])
		if len(cmds) != 3 {
			fmt.Printf(Usage)
			os.Exit(1)
		}
		cli.AddBlock(cmds[2])
	case "printChain":
		fmt.Printf("打印区块链命令被调用\n")
		cli.PrintChain()
	default:
		fmt.Printf("无效命令，请检查 %s\n", cmds[2])
		fmt.Printf(Usage)
		break
	}

}
