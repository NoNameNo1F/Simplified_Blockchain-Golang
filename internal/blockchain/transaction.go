package blockchain

import (
	"fmt"
	"strconv"

	utils "lab1-blockchain-v1.1.0/internal/utils"
)

type Transaction struct {
	Data []byte `json:"data"`
}

func CreateTransaction(data string) *Transaction {
	tx := Transaction{Data: []byte(data)}
	return &tx
}

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
