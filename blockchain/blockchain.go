package blockchain

import (
	"github.com/kerimay/go-blockchain/organizeall"
	"log"
	"time"
)

type Blockchain struct {
	db  organizeall.DBaseInterface
	tip []byte
}

func NewBlockchain(db organizeall.DBaseInterface) *Blockchain {
	var bl Block
	if db.BlockchainExists() == false {
		log.Println("There wasn't a previous blockchain before. New blockchain creation begins...")
		genesisBlock := NewGenesisBlock()
		data := bl.EncodeStruct(genesisBlock)
		newTip := genesisBlock.Hash

		db.NewTransaction(newTip, data)
		return &Blockchain{db, newTip}
	} else {
		newTip := db.QueryTip()
		/*var data string
		b.AddBlock(data)*/
		return &Blockchain{db, newTip}
	}
}

func NewGenesisBlock() *Block {
	log.Println("Genesis Block is being created...")
	genesisBlock := NewBlock([]byte("Genesis Block"), []byte{})
	return genesisBlock
}

func (bc *Blockchain) AddBlock(data string) {
	var bl Block
	bc.tip = bc.db.QueryTip()
	block := NewBlock([]byte(data), bc.tip)
	encodedStruct := bl.EncodeStruct(block)
	bc.db.NewTransaction(block.Hash, encodedStruct)
}

func (bc *Blockchain) Iterator() {
	var b Block
	tip := bc.db.QueryTip()

	for {
		byteBlock := bc.db.QueryBlock(tip)
		block, err := b.DecodeStruct(byteBlock)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("The hash %x belongs to the block...\n Hash: %x\n PrevHash: %x\n Data: %s\n Timestamp: %v\n Nonce: %v\n\n", tip, block.Hash, block.PrevHash, block.Data, block.Timestamp, block.Nonce)
		tip = block.PrevHash
		time.Sleep(time.Second * 2)
		if string(block.Data) == "Genesis Block" {
			break
		}
	}
}
