package bchain_dbase_interaction

import "github.com/kerimay/go-blockchain/blockchain"

type Bchain interface {
	NewBlock(data []byte, prevHash []byte) *blockchain.Block
	BringBlockHash() []byte
}
