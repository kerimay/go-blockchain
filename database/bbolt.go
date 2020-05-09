package database

import (
	"bytes"
	"fmt"
	"github.com/kerimay/go-blockchain/blockchain"
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

type DataBase struct {
	db *bolt.DB
}

var network bytes.Buffer // Stand-in for the network.

func (d *DataBase) OpenDataBase(fileName string) *DataBase {
	log.Println("Connecting the database...")
	db, err := bolt.Open(fileName, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return &DataBase{db}
}

func (d *DataBase) CreateBlocksBucket() {
	d.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("blocks"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func (d *DataBase) NewTransaction(block *blockchain.Block) {
	// Create several keys in a transaction.
	tx, err := d.db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}

	b := tx.Bucket([]byte("blocks"))
	EncodeStruct(block)

	if err = b.Put(block.Hash, network.Bytes()); err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func (d *DataBase) QueryDB() {
	// Iterate over the values in sorted key order.
	tx, err := d.db.Begin(false)
	if err != nil {
		log.Fatal(err)
	}

	var decBlock blockchain.Block
	DecodeStruct(decBlock)

	c := tx.Bucket([]byte("blocks")).Cursor()
	for k, v := c.First(); k != nil; k, v = c.Next() {
		fmt.Printf("%x hash belongs to the block: %s\n", k, v)
	}

	if err = tx.Rollback(); err != nil {
		log.Fatal(err)
	}

	if err = d.db.Close(); err != nil {
		log.Fatal(err)
	}
}
