package database

import (
	bolt "go.etcd.io/bbolt"
	"log"
	"time"
)

type DataBase struct {
}

func (d *DataBase) OpenDataBase() {
	log.Println("Opening the database...")
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
