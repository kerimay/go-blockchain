package blockchain

import "fmt"

type Blockchain struct {
	Blocks []Block
}

var newBlockchain []Block

func CreateBlockchain() Blockchain {
	genesisBlock := GenesisBlock()
	newBlockchain = append(newBlockchain, genesisBlock)
	fmt.Println(Blockchain{Blocks: newBlockchain})

	return Blockchain{newBlockchain}
}

func (b *Block) AddBlockToBlockchain(newBlock Block) Blockchain {
	newBlock = CreateBlock(b.Data)
	newBlockchain = append(newBlockchain, newBlock)
	fmt.Println(Blockchain{Blocks: newBlockchain})
	return Blockchain{newBlockchain}
}

func QueryBlockchain() Blockchain {
	fmt.Println(Blockchain{newBlockchain})
	return Blockchain{newBlockchain}
}
