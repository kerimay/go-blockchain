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

func (bl *Block) NewBlock(data []byte, prevHash []byte) *Block {
	block := &Block{
		Data:      data,
		PrevHash:  prevHash,
		Timestamp: time.Now().Unix(),
	}
	pow := NewProofOfWork(block)
	b := pow.block
	b.Hash, b.Nonce = pow.findHash()
	// ERROR
	//DBase.NewTransaction()
	return block
}

func (bl *Block) BringBlockHash() []byte {
	return bl.Hash // yeterli mi?
}
