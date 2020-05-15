package blockchain

import (
	"github.com/kerimay/go-blockchain/organizeall"
	"log"
)

type Blockchain struct {
	db  organizeall.DBaseInterface
	Tip []byte
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
	bc.Tip = bc.db.QueryTip()
	block := NewBlock([]byte(data), bc.Tip)
	encodedStruct := bl.EncodeStruct(block)
	bc.db.NewTransaction(block.Hash, encodedStruct)
	log.Println("Success!")
	log.Println()
}

func (bc *Blockchain) Iterator() {
	iter := NewBlockchainIterator(bc)

	if iter != nil {
		for iter.hasNext() {
			queryResult := iter.next()
			newPow := NewProofOfWork(queryResult)
			powBool := newPow.isPoWProven(queryResult.Nonce)

			log.Printf("Hash: %x\n PrevHash: %x\n Data: %s\n Timestamp: %v\n Nonce: %v\n PoW: %v\n\n", queryResult.Hash, queryResult.PrevHash, queryResult.Data, queryResult.Timestamp, queryResult.Nonce, powBool)
		}
	}
}
