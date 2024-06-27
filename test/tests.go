package test

import (
	"fmt"

	bc "lab1-blockchain-v1.1.0/internal/blockchain"
)

func TestAddTransaction(transactions []*bc.Transaction, numberOfTrans int) []*bc.Transaction {
	var msg string
	for i := 1; i <= numberOfTrans; i++ {
		msg = fmt.Sprintf("Test Transaction %d", i)
		tx := bc.CreateTransaction(msg)
		transactions = append(transactions, tx)
	}

	bc.ViewTransactions(transactions)

	return transactions
}

func TestAddBlock(blockchain *bc.Blockchain, transactions []*bc.Transaction) (*bc.Blockchain, []*bc.Transaction) {
	if len(transactions) == 0 {
		transactions = TestAddTransaction(transactions, 2)
	}
	blockchain.AddBlock(transactions)
	transactions = nil

	blockchain.ViewBlockchain()

	return blockchain, transactions
}
