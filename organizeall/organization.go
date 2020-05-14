package organizeall

type DBaseInterface interface {
	NewTransaction(hash []byte, data []byte)
	QueryTip() []byte
	QueryBlock(hash []byte) []byte
	BlockchainExists() bool
}
