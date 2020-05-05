package blockchain

import (
	"bytes"
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
	var header []byte
	var nonce int
	targetInt := targetBigInt()

	for nonce = 0; nonce < maxNonce; nonce++ {
		header = hashData(p.prepareData(nonce))
		intHash := hashBigInt(header)

		if intHash.Cmp(targetInt) == -1 { // for making our hash SMALLER than the constant target
			break
		}
	}
	return header, nonce
}

func (p *ProofOfWork) isPoWProven(nonce int) bool {
	targetInt := targetBigInt()
	hash := hashData(p.prepareData(nonce))
	intHash := hashBigInt(hash)
	isValid := intHash.Cmp(targetInt) == -1
	return isValid
}

func (p *ProofOfWork) prepareData(nonce int) []byte {
	b := p.block
	byteTime := intToBytes(b.Timestamp)
	byteNonce := intToBytes(int64(nonce))
	return bytes.Join([][]byte{b.Data, b.PrevHash, byteTime, byteNonce}, []byte{})
}