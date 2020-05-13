package organizeall

type DBaseInterface interface {
	NewTransaction(hash []byte, data []byte)
	QueryTip() []byte
	QueryDB()
	BlockchainExists() bool
}
