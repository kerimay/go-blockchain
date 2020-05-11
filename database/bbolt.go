package database

import (
	"bytes"
	"fmt"
	bchain_dbase_interaction "github.com/kerimay/go-blockchain/bchain-dbase-interaction"
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

type DataBase struct {
	db *bolt.DB
	bchain_dbase_interaction.Bchain
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
	return &DataBase{db, d.Bchain} // doğru mu?
}

func (d *DataBase) NewTransaction(o *bchain_dbase_interaction.Bchain) {
	// Create several keys in a transaction.
	tx, err := d.db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}

	b, err := tx.CreateBucketIfNotExists([]byte("blocks"))
	if err != nil {
		log.Fatal("bucket creation", err)
	}
	EncodeStruct(o)

	if err = b.Put(d.Bchain.BringBlockHash(), network.Bytes()); err != nil {
		log.Fatal(err)
	}

	if err = b.Put([]byte("l"), d.Bchain.BringBlockHash()); err != nil {
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

func (d *DataBase) QueryDB() { // düzenle
	// Iterate over the values in sorted key order.
	tx, err := d.db.Begin(false)
	if err != nil {
		log.Fatal(err)
	}

	var decBlock bchain_dbase_interaction.Bchain
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

// Will be used if needed
func (d *DataBase) IsBlockchain() bool {
	tx, err := d.db.Begin(false)
	if err != nil {
		log.Fatal(err)
	}

	isFull := tx.Bucket([]byte("blocks"))
	if isFull == nil {
		return false
	} else {
		return true
	}
}
