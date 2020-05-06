package main

import (
	"github.com/kerimay/go-blockchain/blockchain"
	"github.com/kerimay/go-blockchain/database"
)

/*type BlockchainDBStructure interface {

}*/

func main() {

	db := database.DataBase{}
	db.OpenDataBase()

	bc := blockchain.CreateBlockchain()
	bc.AddBlock([]byte("Send 2 BTC to Selçuk"))
	bc.AddBlock([]byte("Send 2 BTC to Selçuk"))
	bc.AddBlock([]byte("Send 1 BTC to Cengiz"))
	bc.QueryBlockchain()
}
