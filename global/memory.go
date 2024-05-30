package global

type Memory struct {
	// Memory is a struct that holds the memory of the Ethereum Virtual Machine.
	// List of addresses that have been observed.
	Addresses []string
	// List of transactions that have been observed.
	Transactions []Transaction
}

type Transaction struct {
	// Transaction is a struct that holds the details of a transaction.
	// The address of the sender.
	From string
	// The address of the receiver.
	To string
	// The amount of Ether transferred.
	Value int
}
