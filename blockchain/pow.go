package blockchain

import (
	"bytes"
	"fmt"
	"math"
)

type ProofOfWork struct {
	block *Block
}

func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{block: block}
}


const maxNonce = math.MaxInt64


func (p *ProofOfWork) findHash() ([]byte, int) {
	b := p.block

	var header []byte
	var nonce int
	targetInt := targetBigInt()
	header = p.prepareData()
	hashedHeader := hashData(header)
	for nonce = 0; nonce < maxNonce; nonce++ {
		intHash := hashBigInt(hashedHeader)
		fmt.Println("_________")
		fmt.Printf("Hex hash: %s\n", hashedHeader[:])
		if intHash.Cmp(targetInt) == -1 { // for making our hash SMALLER than the constant target
			b.Hash = hashedHeader
			break
		}
		fmt.Printf("Big Int hash and nonce: %v\n%v\n", intHash, b.Nonce)
		fmt.Printf("target big int: %v\n", targetInt)

	}
	fmt.Println(b.Hash, nonce)
	return b.Hash, nonce
}

func (p *ProofOfWork) isPoWProven() bool {
	/*generatedHash, generatedNonce := findHash()
	if generatedHash < targetBigInt() {

	}*/
	return true
}

func (p *ProofOfWork) prepareData() []byte {
	b := p.block
	byteTime := intToBytes(b.Timestamp)
	byteNonce := intToBytes(int64(b.Nonce))
	return bytes.Join([][]byte{b.Data, b.PrevHash, byteTime, byteNonce}, []byte{})
}

