package database

import (
	"fmt"
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
	return &DataBase{db} // doğru mu?
}

func (d *DataBase) NewTransaction(hash []byte, data []byte) {
	log.Println("Transaction process begins...")
	d.db.Update(func(tx *bolt.Tx) error {
		log.Println("Bucket creation begins...")
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
		log.Println("Put the last hash.")
		return nil
	})
}

func (d *DataBase) QueryTip() []byte {
	var tip []byte
	d.db.View(func(tx *bolt.Tx) error {
		log.Println("Tip is being called...")
		tip = tx.Bucket([]byte(blockBucket)).Get([]byte("l"))
		return nil
	})
	return tip
}

func (d *DataBase) QueryDB() { // düzenle
	d.db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte(blockBucket)).Cursor()
		log.Println("Database is being queried...")
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("%x hash belongs to the block: %s\n", k, v)
		}
		return nil
	})
}

func (d *DataBase) BlockchainExists() bool {
	var isFull *bolt.Bucket
	d.db.View(func(tx *bolt.Tx) error {
		isFull = tx.Bucket([]byte(blockBucket))
		return nil
	})
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
