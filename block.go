package main

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Version      uint64 //区块版本号
	PreBlockHash []byte //前区块hash
	MerkleRoot   []byte
	TimeStamp    uint64 //从1970。1。1至今的秒数
	Difficulity  uint64 //难度值
	Nonce        uint64 //随机数，挖矿找的就是它
	Data         []byte //数据体
	Hash         []byte //当前区块哈希，原本不存在，为了方便才加进来的
}

func NewBlock(data string, preBlockHash []byte) *Block {
	block := Block{
		Version:      00,
		PreBlockHash: preBlockHash,
		MerkleRoot:   []byte{},
		TimeStamp:    uint64(time.Now().Unix()),
		Difficulity:  10,
		Nonce:        10,
		Data:         []byte(data),
		Hash:         []byte{},
	}
	//block.SetHash()
	pow := NewProofOfWork(&block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return &block
}

func (block *Block) SetHash() {
	temp := [][]byte{unit2Byte(block.Version),
		block.PreBlockHash,
		block.MerkleRoot,
		unit2Byte(block.TimeStamp),
		unit2Byte(block.Difficulity),
		unit2Byte(block.Nonce),
		block.Data,
	}
	data := bytes.Join(temp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}
