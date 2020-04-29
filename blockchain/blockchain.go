package blockchain

import (
	"fmt"
)

type Blockchain struct {
	Blocks []*Block
}

func CreateBlockchain() *Blockchain {
	genesisBlock := NewBlock([]byte("Genesis Block"), []byte{})

	return &Blockchain{[]*Block{genesisBlock}}
}

func (bc *Blockchain) AddBlock(data []byte)  {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) QueryBlockchain() {
	for _, b := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("PoW: %v\n", isPoWProven())
		fmt.Printf("\n")
	}
}