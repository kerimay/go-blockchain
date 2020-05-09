package blockchain

import (
	"fmt"
	"github.com/kerimay/go-blockchain/database"
)

type Blockchain struct {
	Blocks []*Block
	*database.DataBase
}

var DBase database.DataBase

func CreateBlockchain(fileName string) *Blockchain {

	db := DBase.OpenDataBase(fileName)
	if db == nil {
		DBase.CreateBlocksBucket()
	}
	genesisBlock := NewBlock([]byte("Genesis Block"), []byte{})

	return &Blockchain{[]*Block{genesisBlock}, db}
}

func (bc *Blockchain) AddBlock(data []byte) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) QueryBlockchain() {

	for _, b := range bc.Blocks {
		p := NewProofOfWork(b)
		fmt.Printf("PrevHash: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("PoW: %v\n", p.isPoWProven(b.Nonce))
		fmt.Printf("\n")
	}
}
