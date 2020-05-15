package main

import (
	"github.com/kerimay/go-blockchain/blockchain"
	"github.com/kerimay/go-blockchain/database"
	"log"
)

const dbFile = "blockchain.db"

func main() {
	db := database.NewDataBase(dbFile)
	bc := blockchain.NewBlockchain(db)
	cli := NewCLI(bc)
	defer func() error {
		err := db.Close()
		if err != nil {
			log.Fatal("Database couldn't be closed", err)
			return err
		}
		return nil
	}()

	cli.Run()
}
