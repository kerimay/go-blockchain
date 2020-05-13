package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	Data      []byte
	Hash      []byte
	PrevHash  []byte
	Timestamp int64
	Nonce     int
}

func NewBlock(data []byte, prevHash []byte) *Block {
	block := &Block{
		Data:      data,
		PrevHash:  prevHash,
		Timestamp: time.Now().Unix(),
	}
	pow := NewProofOfWork(block)
	b := pow.block
	b.Hash, b.Nonce = pow.findHash()
	return block
}

func (bl *Block) EncodeStruct(block *Block) []byte {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(&block)
	if err != nil {
		log.Fatal("encode:", err)
	}
	return network.Bytes()
}

func (bl *Block) DecodeStruct(data []byte) error {
	var network bytes.Buffer //Reader?
	x := bytes.NewReader(data)
	dec := gob.NewDecoder(&network)
	err := dec.Decode(x)
	if err != nil {
		log.Fatal("decode err", err)
	}
	return err
}
