package utils

import (
	"fmt"
)

// ShowHelpCommand displays the help commands for the application.
//
// It prints a table with the available commands and their descriptions.
// The commands are categorized into features, helpers, and testing.
// Each command is represented by a number and a description.
// The function does not take any parameters and does not return anything.
func ShowHelpCommand() {
	fmt.Println("|------------------------------------ Help Commands -------------------------------------|")
	fmt.Println("|[FEATURE]1. View_Blockchain: Displays information about the entire blockchain.          |")
	fmt.Println("|[FEATURE]2. Add_Transaction {message}: Adds a new transaction to the transaction list.  |")
	fmt.Println("|[FEATURE]3. View_Transactions: Shows transactions need to be added into new block.      |")
	fmt.Println("|[FEATURE]4. Add_Block: Adds current list transactions into New Block.                   |")
	fmt.Println("|[FEATURE]5. Verify_Block: Verifies the block is valid or not.                           |")
	fmt.Println("|[FEATURE]6. Print_MerkleTree: View Merkle Tree of specific block.                       |")
	fmt.Println("|[HELPERS]7. Help: Displays this help message                                            |")
	fmt.Println("|[EXITING]8. Exit: Exits the application                                                 |")
	fmt.Println("|----------------------------------------------------------------------------------------|")
	fmt.Println("|[TESTING] 1. Test1: Test1 {numberOfTransactions}: Add {n} transactions to test          |")
	fmt.Println("|[TESTING] 2. Test2: Test2: Testing add Block into Blockchain                            |")
	fmt.Println("|----------------------------------------------------------------------------------------|")
}

// Authors: https://github.com/NoNameNo1F/Simplified_Blockchain-Golang
