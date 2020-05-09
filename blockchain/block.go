package blockchain

import (
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

	DBase.NewTransaction(block)
	return block
}
