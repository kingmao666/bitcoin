package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func unit2Byte(num uint64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buffer.Bytes()
}
