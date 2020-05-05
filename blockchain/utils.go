package blockchain

import (
	"crypto/sha256"
	"math/big"
	"strconv"
)

const targetBits = 24

func intToBytes(data int64) []byte {
	b := []byte(strconv.FormatInt(data, 10))
	return b
}

func targetBigInt() *big.Int {
	targetBig := big.NewInt(1)
	targetBig.Lsh(targetBig, 256 - targetBits)
	return targetBig
}

func hashBigInt(header []byte) *big.Int {
	z := new(big.Int)
	bigHash := z.SetBytes(header[:])
	return bigHash
}

func hashData(data []byte) []byte {
	header := sha256.Sum256(data)
	return header[:]
}
