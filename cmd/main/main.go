package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	bc "lab1-blockchain-v1.1.0/internal/blockchain"
	"lab1-blockchain-v1.1.0/internal/utils"
	test "lab1-blockchain-v1.1.0/test"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	blockchain := bc.LoadBlockchain()

	blockchainChannel := blockchain.FetchingDatabase()

	var transactions []*bc.Transaction

	// Run in a separate goroutine
	go func() {
		for {
			fmt.Print("HCMUS:\\Users\\20127670> ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			parts := strings.Fields(input)

			if len(parts) == 0 {
				continue
			}

			command := strings.ToLower(parts[0])

			switch command {
			case "view_blockchain":
				blockchain.ViewBlockchain()
			// logging for view blockchain

			case "add_transaction":
				if len(parts) < 2 {
					fmt.Println("Usage: add_transaction {message}")
					continue
				}
				message := strings.Join(parts[1:], " ")
				newTx := bc.CreateTransaction(message)
				transactions = append(transactions, newTx)
				// logging for add new transaction

			case "add_block":
				if len(transactions) == 0 {
					fmt.Printf("There are no transactions to add\n")
					continue
				}

				blockchain.AddBlock(transactions)
				transactions = nil
				blockchain.SaveBlockchainToFile()
			case "verify_block":
				if len(parts) < 2 {
					fmt.Println("Usage: verify_block {hash of merkle_root}")
					continue
				}
				data := strings.Join(parts[1:], " ")

				blockchain.VerifyBlock(data)
				// logging for verify status

			case "view_transactions":
				bc.ViewTransactions(transactions)
				// logging for view current transactions not added yet

			case "print_merkletree":
				if len(parts) < 2 {
					fmt.Println("Usage: print_merkletree {hash of merkle_root}")
					continue
				}

				data := strings.Join(parts[1:], " ")

				blockchain.ViewMerkleTree(data)
				//logging for blockchain reset

			case "help":
				utils.ShowHelpCommand()
			case "test1":
				transactions = test.TestAddTransaction(transactions, 2)
			case "test2":
				blockchain, transactions = test.TestAddBlock(blockchain, transactions)
			case "exit":
				fmt.Println("Exiting...")
				blockchain.SaveBlockchainToFile()
				os.Exit(0)

			default:
				fmt.Println("Unknown command. Type 'Help' for instructions.")
			}
		}
	}()

	for {
		select {
		case newBlockchain := <-blockchainChannel:
			blockchain = newBlockchain
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}
