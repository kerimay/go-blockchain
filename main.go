package main

import (
	"go-bootcamp/blockchain"
)

func main() {
	bc := blockchain.CreateBlockchain()
	bc.AddBlock([]byte("Send 2 BTC to Selçuk"))
	bc.AddBlock([]byte("Send 2 BTC to Selçuk"))
	bc.AddBlock([]byte("Send 1 BTC to Cengiz"))
	bc.QueryBlockchain()

}
