package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "add|delete|search")
		os.Exit(1)
	}

	if !awsAccountCheck() {
		fmt.Println("Could not authenticate, exiting...")
		os.Exit(2)
	}

	mongo := initMongo()

	switch os.Args[1] {
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter account name or site URL: ")
		accountName, _ := reader.ReadString('\n')
		fmt.Print("Enter password for account: ")
		accountPassword, _ := reader.ReadString('\n')
		accountName = strings.Trim(accountName, "\n")
		accountPassword = strings.Trim(accountPassword, "\n")

		addAccount(accountName, accountPassword, mongo)
	case "delete":
		fmt.Println("Delete mode...")

	case "search":
		fmt.Println("Search mode...")

	default:
		fmt.Println("Unrecognized command...")
	}
}
