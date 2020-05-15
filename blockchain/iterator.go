package blockchain

import (
	"github.com/kerimay/go-blockchain/organizeall"
	"log"
)

type StatefulIterator interface {
	next() []byte
	hasNext() bool
}

type Iterator struct {
	Hash []byte
	db   organizeall.DBaseInterface
}

func NewBlockchainIterator(bc *Blockchain) *Iterator {
	return &Iterator{bc.Tip, bc.db}
}

func (it *Iterator) next() *Block {
	var b Block
	byteBlock := it.db.QueryBlock(it.Hash)
	block, err := b.DecodeStruct(byteBlock)
	if err != nil {
		log.Fatal("decoding", err)
	}
	it.Hash = block.PrevHash
	return block
}
func (it *Iterator) hasNext() bool {
	return len(it.Hash) > 0
}
