package database

import (
	"bytes"
	"encoding/gob"
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

func (d *DataBase) RunDatabase(data []byte, prevHash []byte) {
	d.OpenDataBase()
	d.CreateBlocksBucket() // is there a blockchain?
	d.NewTransaction(data, prevHash)
}

func (d *DataBase) OpenDataBase() {
	log.Println("Connecting the database...")
	db, err := bolt.Open("blockchain.db", 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
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

func (d *DataBase) NewTransaction(data []byte, prevHash []byte) {
	// Create several keys in a transaction.
	tx, err := d.db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}

	b := tx.Bucket([]byte("blocks"))

	block := blockchain.NewBlock(data, prevHash)
	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err = enc.Encode(&block)
	if err != nil {
		log.Fatal("encode:", err)
	}

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
	dec := gob.NewDecoder(&network)
	err = dec.Decode(&decBlock)
	if err != nil {
		log.Fatal("decode err", err)
	}
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
