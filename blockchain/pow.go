package blockchain

import (
	"fmt"
	"math/big"
)

const target string = "0000010000000000000000000000000000000000000000000000000000000000"
var b Block

func findHash() ([]byte, int) {
	b.Nonce = 0
	for  {
		if hashBigInt().Cmp(targetBigInt()) != -1 { // for making our hash SMALLER than the constant target
			continue
		}
		b.Nonce++
	}
	fmt.Println(b.Hash, b.Nonce)
	return b.Hash, b.Nonce
}

func isPoWProven() bool {

	return true
}

func targetBigInt() *big.Int {
	targetBig, _ := new(big.Int).SetString(target, 10)
	return targetBig
}

func hashBigInt()  *big.Int {
	hashedHeader := b.hashedHeader()
	z := new(big.Int)
	intHash := z.SetBytes(hashedHeader[:])
	return intHash
}