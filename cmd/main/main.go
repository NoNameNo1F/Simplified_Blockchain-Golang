package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	bc "lab1-blockchain-v1.1.0/internal/blockchain"
	"lab1-blockchain-v1.1.0/internal/logging"
	"lab1-blockchain-v1.1.0/internal/utils"
	test "lab1-blockchain-v1.1.0/test"
)

// main is the entry point of the program where it interacts with the blockchain through user commands.
//
// It reads user input, processes commands to interact with the blockchain, and handles various functionalities like adding transactions, blocks, verifying blocks, viewing transactions, and displaying help.
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

				logging.SetLogOutput(false, true)
				logging.Log("INFO", "Viewed blockchain")

			case "add_transaction":
				if len(parts) < 2 {
					logging.SetLogOutput(true, true)
					logging.Log("ERROR", "Failed to add transaction: insufficient arguments")
					fmt.Println("Usage: add_transaction {message}\n")
					continue
				}
				message := strings.Join(parts[1:], " ")
				newTx := bc.CreateTransaction(message)
				transactions = append(transactions, newTx)

				logging.SetLogOutput(false, true)
				logging.Log("INFO", fmt.Sprintf("Added new transaction: %s", message))

			case "add_block":
				if len(transactions) == 0 {
					logging.SetLogOutput(true, true)
					logging.Log("ERROR", "Failed to add block: there are no transactions")
					continue
				}

				blockchain.AddBlock(transactions)
				transactions = nil
				blockchain.SaveBlockchainToFile()

				logging.SetLogOutput(false, true)
				logging.Log("INFO", "Added new block to blockchain")

			case "verify_block":
				if len(parts) < 2 {
					logging.SetLogOutput(true, true)
					logging.Log("ERROR", "Failed to verify block: insufficient arguments")
					fmt.Println()
					fmt.Println("Usage: verify_block {hash of merkle_root}\n")
					continue
				}

				data := strings.Join(parts[1:], " ")
				blockchain.VerifyBlock(data)

				logging.SetLogOutput(false, true)
				logging.Log("INFO", fmt.Sprintf("Verified block with merkle root: %s", data))

			case "view_transactions":
				bc.ViewTransactions(transactions)
				logging.SetLogOutput(false, true)
				logging.Log("INFO", "Viewed current transactions not added yet")

			case "print_merkletree":
				if len(parts) < 2 {
					fmt.Println("Usage: print_merkletree {hash of merkle_root}\n")
					logging.Log("ERROR", "Failed to print merkle tree: insufficient arguments")
					continue
				}

				data := strings.Join(parts[1:], " ")
				blockchain.ViewMerkleTree(data)

				logging.SetLogOutput(false, true)
				logging.Log("INFO", fmt.Sprintf("Printed merkle tree for block with merkle root: %s", data))

			case "help":
				utils.ShowHelpCommand()
				logging.SetLogOutput(false, true)
				logging.Log("INFO", "Displayed help commands")

			case "test1":
				transactions = test.TestAddTransaction(transactions, 2)
				logging.SetLogOutput(false, true)
				logging.Log("INFO", "Ran test1 to add transactions")

			case "test2":
				blockchain, transactions = test.TestAddBlock(blockchain, transactions)
				logging.SetLogOutput(false, true)
				logging.Log("INFO", "Ran test2 to add block")

			case "exit":
				blockchain.SaveBlockchainToFile()
				logging.SetLogOutput(true, true)
				logging.Log("INFO", "Exiting application...")
				os.Exit(0)

			default:
				logging.SetLogOutput(true, true)
				logging.Log("ERROR", fmt.Sprintf("Unknown command: %s.", command))
				fmt.Print("Type 'Help' for instructions.\n")
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

// Authors: https://github.com/NoNameNo1F/Simplified_Blockchain-Golang
