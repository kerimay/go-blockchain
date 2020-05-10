package main

import (
	"github.com/kerimay/go-blockchain/blockchain"
)

const dbFile = "blockchain.db"

func main() {

	blockchain.RunBlockchain(dbFile, []byte("Send 2 BTC to Selçuk"))
	blockchain.RunBlockchain(dbFile, []byte("Send 2 BTC to Selçuk"))
	blockchain.RunBlockchain(dbFile, []byte("Send 1 BTC to Cengiz"))

	//blockchain.Blockchain.QueryBlockchain()
}
