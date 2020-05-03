package blockchain

import (

	"fmt"
	"math/big"
)

type ProofOfWork struct {
	*Block
	*Target
}

type Target struct {
	targetBits int
}

const maxNonce = 4294967295
var b Block // this will be removed

func findHash() ([]byte, int) {

	targetInt := targetBigInt()
	for b.Nonce = 0; b.Nonce < maxNonce; b.Nonce++ {
		intHash, byteHash := hashBigInt()
		fmt.Printf("%s\n", byteHash[:])
		if intHash.Cmp(targetInt) == -1 { // for making our hash SMALLER than the constant target
			break
		}
		fmt.Println(intHash, b.Nonce)
		fmt.Println(targetInt)
		fmt.Println((*big.Int).BitLen(targetInt))
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

	targetBig := big.NewInt(1)
	targetBig.Lsh(targetBig, 232)

	return targetBig
}

func hashBigInt()  (*big.Int, string) {
	hashedHeader := b.hashedHeader()
	strHashedHeader := fmt.Sprintf("%x", hashedHeader)
	z := new(big.Int)
	intHash := z.SetBytes(hashedHeader[:])
	return intHash, strHashedHeader
}