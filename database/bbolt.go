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
var tip []byte

func (d *DataBase) OpenDataBase(fileName string) *DataBase {
	log.Println("Connecting the database...")
	db, err := bolt.Open(fileName, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return &DataBase{db}
}

func (d *DataBase) NewTransaction(block *blockchain.Block) {
	// Create several keys in a transaction.
	tx, err := d.db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}

	b, err := tx.CreateBucketIfNotExists([]byte("buckets"))
	if err != nil {
		log.Fatal("bucket creation", err)
	}
	EncodeStruct(block)

	if err = b.Put(block.Hash, network.Bytes()); err != nil {
		log.Fatal(err)
	}

	if err = b.Put([]byte("l"), block.Hash); err != nil {
		log.Fatal(err)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func (d *DataBase) QueryTip() []byte {
	b, err := d.db.Begin(false)
	if err != nil {
		log.Fatal(err)
	}
	tip = b.Bucket([]byte("blocks")).Get([]byte("l"))
	return tip
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
