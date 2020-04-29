package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Data []byte
	Hash []byte
	PrevHash []byte
	Timestamp int64
	Nonce int
}

func NewBlock(data []byte, prevHash []byte) *Block {

	block := &Block{
		Data: data,
		PrevHash: prevHash,
		Timestamp: time.Now().Unix(),
	}
	block.Hash, block.Nonce = findHash()
	return block
}

func (b *Block) hashedHeader() []byte {

	byteTime := []byte(strconv.FormatInt(b.Timestamp, 10))

	headers := bytes.Join([][]byte{b.Data, b.PrevHash, byteTime}, []byte{})
	hashedHeader := sha256.Sum256(headers)


	return hashedHeader[:]
}