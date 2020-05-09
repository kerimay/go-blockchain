package database

import (
	"encoding/gob"
	"log"
)

func EncodeStruct(x interface{}) {
	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(&x)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

func DecodeStruct(x interface{}) {
	dec := gob.NewDecoder(&network)
	err := dec.Decode(&x)
	if err != nil {
		log.Fatal("decode err", err)
	}
}
