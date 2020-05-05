package blockchain

import (
	"crypto/sha256"
	"math/big"
	"strconv"
)

const targetBits = 24

func targetBigInt() *big.Int {

	targetBig := big.NewInt(1)
	targetBig.Lsh(targetBig, 256 - targetBits)

	return targetBig
}

func hashBigInt(header []byte)  *big.Int {
	hashedHeader := hashData(header)
	z := new(big.Int)
	intHash := z.SetBytes(hashedHeader)
	return intHash
}

func intToBytes(data int64) []byte {
	return []byte(strconv.FormatInt(data, 10))
}

func hashData(headers []byte) []byte {
	hashedHeader := sha256.Sum256(headers)
	return hashedHeader[:]
}