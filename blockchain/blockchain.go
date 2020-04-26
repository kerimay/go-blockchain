package blockchain

type Blockchain struct {
	blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (bc *Blockchain) Blocks() []*Block {
	return bc.blocks
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
