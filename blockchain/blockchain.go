package blockchain

import (
	"fmt"
	"github.com/kerimay/go-blockchain/database"
	"log"
)

type Blockchain struct {
	Blocks []*Block
	*database.DataBase
}

var DBase *database.DataBase
var tip = DBase.QueryTip()
var bl *Block

func RunBlockchain(fileName string, data []byte) {
	db := DBase.OpenDataBase(fileName)
	if db.IsBlockchain() == false {
		CreateBlockchain() // bağlı mı olmalı?
	}
	/*if db == nil { // ?????
	}*/
	b := &Blockchain{Blocks: []*Block{}} // ??????????
	b.AddBlock(data)
}

func CreateBlockchain() *Blockchain {
	genesisBlock := bl.NewBlock([]byte("Genesis Block"), []byte{})
	tip = genesisBlock.Hash // gerek var mı?
	return &Blockchain{[]*Block{genesisBlock}, DBase}
}

func (bc *Blockchain) AddBlock(data []byte) {
	newBlock := bl.NewBlock(data, tip)      // tip çakışma yaşıyor mu
	bc.Blocks = append(bc.Blocks, newBlock) // bunlara gerek kalmayacak

	if string(tip) != string(newBlock.Hash) {
		log.Fatal("tip and hash are not equal")
	}
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
