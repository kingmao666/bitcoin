package main

//创建区块链，使用Block数组模拟
type BlockChain struct {
	Blocks []*Block
}

//创建区块链
func NewBlockChain() *BlockChain {
	gen := NewBlock("The Times 03/Jan/2009 Chancellor on brink of second bailout for banks", []byte{0x0000000000000000})
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
