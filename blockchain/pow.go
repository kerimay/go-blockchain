package blockchain

import (
	"fmt"
	"math/big"
)

const target string = "0000010000000000000000000000000000000000000000000000000000000000"
const maxNonce = 4294967295
var b Block

func findHash() ([]byte, int) {

	targetInt := targetBigInt()
	for b.Nonce = 0; b.Nonce < maxNonce; b.Nonce++ {
		intHash, byteHash := hashBigInt()
		if intHash.Cmp(targetInt) == -1 { // for making our hash SMALLER than the constant target
			break
		}
		fmt.Println(byteHash, b.Nonce)
	}
	fmt.Println(b.Hash, b.Nonce)
	return b.Hash, b.Nonce
}

func isPoWProven() bool {
	/*generatedHash, generatedNonce := findHash()
	if generatedHash < targetBigInt() {

	}*/
	return true
}

func targetBigInt() *big.Int {
	targetBig, _ := new(big.Int).SetString(target, 10)
	return targetBig
}

func hashBigInt()  (*big.Int, string) {
	hashedHeader := b.hashedHeader()
	strHashedHeader := fmt.Sprintf("%x", hashedHeader)
	z := new(big.Int)
	intHash := z.SetBytes(hashedHeader[:])
	return intHash, strHashedHeader
}