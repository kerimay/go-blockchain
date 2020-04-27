package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Data string
	Hash string
	PrevHash string
	Timestamp string
}

func GenesisBlock() Block {
	newHash := sha256.Sum256([]byte("Genesis Block"))
	newStrHash := fmt.Sprintf("%x", newHash)
	genesisTimestamp := time.Now().UTC().String()

	genesisBlock := Block{"Genesis Block",newStrHash, "", genesisTimestamp}
	return genesisBlock
}

func CreateBlock(data string) Block {

	var newPrevHash string
	newPrevHash = newBlockchain[len(newBlockchain)-1].Hash

	newHash := sha256.Sum256([]byte(data))
	newStrHash := fmt.Sprintf("%x", newHash)

	newTimestamp := time.Now().UTC().String()

	returnedBlock :=  Block{data, newStrHash, newPrevHash , newTimestamp}
	fmt.Println(returnedBlock)

	return returnedBlock
}

