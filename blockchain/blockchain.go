package blockchain

import (
	"github.com/kerimay/go-blockchain/organizeall"
	"log"
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

func (bc *Blockchain) QueryBlockchain() {
	/*for _, b := range bc.Blocks {
		p := NewProofOfWork(b)
		fmt.Printf("PrevHash: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("PoW: %v\n", p.isPoWProven(b.Nonce))
		fmt.Printf("\n")
	}*/
}
