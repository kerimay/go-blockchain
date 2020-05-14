package database

import (
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

type DataBase struct {
	db *bolt.DB
}

var blockBucket = "blocks"

func NewDataBase(fileName string) *DataBase {
	log.Println("Connecting the database...")
	db, err := bolt.Open(fileName, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return &DataBase{db}
}

func (d *DataBase) NewTransaction(hash []byte, data []byte) {
	log.Println("Mining process begins...")
	err := d.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(blockBucket))
		if err != nil {
			log.Fatal("bucket creation", err)
		}
		log.Println("Bucket created.")
		if err = b.Put(hash, data); err != nil {
			log.Fatal(err)
		}
		log.Println("Block is attached to the hash.")
		if err = b.Put([]byte("l"), hash); err != nil {
			log.Fatal(err)
		}
		return nil
	})
	if err != nil {
		log.Fatal("Transaction creation control", err)
	}
}

func (d *DataBase) QueryTip() []byte {
	var tip []byte
	err := d.db.View(func(tx *bolt.Tx) error {
		log.Println("Tip is being called...")
		tip = tx.Bucket([]byte(blockBucket)).Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Fatal("Tip control", err)
	}
	return tip
}

func (d *DataBase) QueryBlock(hash []byte) []byte {
	var byteBlock []byte
	err := d.db.View(func(tx *bolt.Tx) error {
		log.Println("Block is being called...")
		buck := tx.Bucket([]byte(blockBucket))
		byteBlock = buck.Get(hash)
		return nil
	})
	if err != nil {
		log.Fatal("Block query", err)
	}
	return byteBlock
}

func (d *DataBase) BlockchainExists() bool {
	var isFull *bolt.Bucket
	err := d.db.View(func(tx *bolt.Tx) error {
		isFull = tx.Bucket([]byte(blockBucket))
		return nil
	})
	if err != nil {
		log.Fatal("Bucket control", err)
	}
	if isFull == nil {
		log.Println("Database is empty...")
		return false
	} else {
		log.Println("Database has a blockchain...")
		return true
	}
}

func (d *DataBase) Close() error {
	return d.db.Close()
}
