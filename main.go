package main

import (
	"github.com/kerimay/go-blockchain/blockchain"
)

func main() {

	/*db := database.DataBase{}
	db.RunDatabase()*/

	bc := blockchain.CreateBlockchain()
	bc.AddBlock([]byte("Send 2 BTC to Selçuk"))
	bc.AddBlock([]byte("Send 2 BTC to Selçuk"))
	bc.AddBlock([]byte("Send 1 BTC to Cengiz"))
	bc.QueryBlockchain()
}
