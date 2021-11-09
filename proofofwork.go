package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block  *Block
	target *big.Int //系统提供的，固定的
}

func NewProofOfWork(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}
	//写难度值，难度值应该是推导出来的，这一版为了简化写成固定值，然后去推导
	// 00001000000000000000000
	//16进制格式字符串
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	var binIntTmp big.Int
	binIntTmp.SetString(targetStr, 16)
	pow.target = &binIntTmp
	return &pow
}

//pow运算函数，同时获取Hash和Nonce
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//1，获取block数据
	//2。拼接nonce
	//3。sha256
	//4。与难度值比较
	//4。1哈希值大于难度值，nonce++
	//4。2哈希值小与难度值，挖矿成功，退出

	var nonce uint64
	var hash [32]byte

	for {
		fmt.Printf("%x\r", hash)
		hash = sha256.Sum256(pow.prepareData(nonce))
		var bigIntTmp big.Int
		bigIntTmp.SetBytes(hash[:])
		if bigIntTmp.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功 nonce  %d,哈希值为 %x\n", nonce, hash)
			//fmt.Printf("挖矿成功！nonce: %d, 哈希值为: %x\n", nonce, hash)
			break
		} else {
			nonce++
		}
	}
	return hash[:], nonce

}

func (pow *ProofOfWork) prepareData(nonce uint64) []byte {
	block := pow.block
	temp := [][]byte{
		unit2Byte(block.Version),
		block.PreBlockHash,
		block.MerkleRoot,
		unit2Byte(block.TimeStamp),
		unit2Byte(block.Difficulity),
		block.Data,
		unit2Byte(nonce),
	}
	data := bytes.Join(temp, []byte{})
	return data
}
