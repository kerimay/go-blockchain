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
		fmt.Println("_________")
		fmt.Printf("Hex hash: %s\n", byteHash[:])
		if intHash.Cmp(targetInt) == -1 { // for making our hash SMALLER than the constant target
			break
		}
		fmt.Printf("Big Int hash and nonce: %v\n%v\n", intHash, b.Nonce)
		fmt.Printf("target big int: %v\n", targetInt)
		fmt.Printf("hash length: %d\n", len(byteHash))
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
	intHash := z.SetBytes(hashedHeader)
	return intHash, strHashedHeader
}