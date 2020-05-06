package database

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

var network bytes.Buffer // Stand-in for the network.

func encodeStruct(x struct{}) {

	// Create an encoder and send a value.
	enc := gob.NewEncoder(&network)
	err := enc.Encode(x)
	if err != nil {
		log.Fatal("encode:", err)
	}
}

func decodeStruct(x struct{}) {
	// Create a decoder and receive a value.
	dec := gob.NewDecoder(&network)
	err := dec.Decode(&x)
	if err != nil {
		log.Fatal("decode:", err)
	}
	fmt.Println(x)
}
