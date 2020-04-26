package main

import (
	"fmt"
	"go-blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC to Selcuk")
	bc.AddBlock("Send 2 more BTC to Selcuk")

	for _, block := range bc.Blocks() {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
