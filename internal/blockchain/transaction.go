package blockchain

import (
	"fmt"
	"strconv"

	utils "lab1-blockchain-v1.1.0/internal/utils"
)

type Transaction struct {
	Data []byte `json:"data"`
}

// CreateTransaction creates a new Transaction object with the given data.
//
// Parameters:
// - data: a string representing the data for the transaction.
//
// Returns:
// - *Transaction: a pointer to the newly created Transaction object.
func CreateTransaction(data string) *Transaction {
	tx := Transaction{Data: []byte(data)}
	return &tx
}

// ViewTransactions prints a list of transactions with their corresponding IDs and data.
//
// Parameters:
// - transactions: a slice of pointers to Transaction structs representing the transactions to be displayed.
//
// Return type: None.
func ViewTransactions(transactions []*Transaction) {
	fmt.Println("|-List of Transactions ------------------------------------------------------------------|")
	size := 0
	spacings := ""
	for index, tx := range transactions {
		size = 90 - 27 - len(strconv.Itoa(index+1)) - len(tx.Data)
		spacings = utils.GenerateSpacing(size)
		fmt.Printf("| Transaction Id: %d, Data: %s%s|\n", index+1, tx.Data, spacings)
	}
	fmt.Printf("|----------------------------------------------------------------------------------------|\n")
}

// Authors: https://github.com/NoNameNo1F/Simplified_Blockchain-Golang
