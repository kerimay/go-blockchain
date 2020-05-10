package blockchain

import (
	"fmt"
	"github.com/kerimay/go-blockchain/database"
	"go-bootcamp/blockchain"
)

type Blockchain struct {
	Blocks []*Block
}

var DBase database.DataBase
var tip = DBase.QueryTip()

func RunBlockchain(fileName string, data []byte) {
	db := DBase.OpenDataBase(fileName)
	if db == nil { // ?????
		blockchain.CreateBlockchain()
	}
	b := &Blockchain{Blocks: []*Block{{Hash: tip}}} // ??????????
	b.AddBlock(data)
}

func (bc *Blockchain) CreateBlockchain() *Blockchain {
	genesisBlock := NewBlock([]byte("Genesis Block"), []byte{})
	tip = genesisBlock.Hash // gerek var mı?
	return &Blockchain{[]*Block{genesisBlock}}
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
