package main

import (
	"go-bootcamp/blockchain"
)

func main() {
	blockchain.CreateBlockchain()
	blockchain.CreateBlock("Send 2 BTC to Selçuk")
	blockchain.CreateBlock("Send 2 BTC to Selçuk")
	blockchain.CreateBlock("Send 1 BTC to Cengiz")
	blockchain.QueryBlockchain()

}
