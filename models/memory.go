package models

type Memory struct {
	// Memory is a struct that holds the memory of the Ethereum Virtual Machine.
	// List of addresses that have been observed.
	Addresses        map[string]bool
	Transactions     map[string][]Transaction
	CurrentBlock     CurrentBlock
	RecordedTxHashes map[string]bool
}

// type Memory map[string]interface{}

type Transaction struct {
	// Transaction is a struct that holds the details of a transaction.
	// The address of the sender.
	From string
	// The address of the receiver.
	To string
	// The amount of Ether transferred.
	Value string
	// The type of transaction.
	TcType string
	// The hash of the transaction.
	Hash string
}

type CurrentBlock struct {
	// The number of the current block.
	Number int
	// The hash of the current block.
	Hash string
}

var MemoryInstance *Memory

func MemoryInitialize() {
	MemoryInstance = &Memory{
		Addresses:        map[string]bool{},
		Transactions:     map[string][]Transaction{},
		CurrentBlock:     CurrentBlock{},
		RecordedTxHashes: map[string]bool{},
	}
}
