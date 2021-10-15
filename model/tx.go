package model

type StoredTxn struct {
	// Name of the project this transaction was for.
	ProjectName string
	//Hash is the hash of the transaction, in hex
	Hash string `gorm:"unique"`
	// Pending is the pending status of the transaction
	Pending bool
	// Status is the final status og the transaction, whether it failed or succeeded
	Status uint
	//Data is the hex encoded version of the data submitted in the  transaction
	Data string
	// Value *big.Int
	// The nonce of the transaction
	Nonce uint64
}
