package model

type Txn struct {
}

type QueuedTxn struct {
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
	// The nonce of the transaction
	Nonce uint64

	// The value of this transaction
	Value string
	// The amount minted for this transaction - used to properly calculate gas.
	Amount uint64
}

type FinalTxn struct {
	// Name of the project this transaction was for.
	ProjectName string
	//Hash is the hash of the transaction, in hex
	Hash string `gorm:"unique"`
	// Status is the final status og the transaction, whether it failed or succeeded
	Status uint

	// The value of this transaction
	Value string
	// The total cost of this transaction
	Cost string
	// the total gas cost
	Gas uint64
	// the gas price
	GasPrice string
}
