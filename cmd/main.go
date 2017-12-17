package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cnerepo/goshush/pkg/goshush"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "add|delete|search")
		os.Exit(1)
	}

	gsSrv := new(goshush.Service)
	gsSrv, err := goshush.Connect()
	if err != nil {
		fmt.Println("Could not configure GoShush Service, exiting...")
		os.Exit(2)
	}

	if err := goshush.Authenticate(); err != nil {
		fmt.Println("Could not authenticate, exiting...")
		os.Exit(2)
	}

	reader := bufio.NewReader(os.Stdin)
	switch os.Args[1] {
	case "add":
		fmt.Print("Enter account name or site URL: ")
		accountName, _ := reader.ReadString('\n')
		fmt.Print("Enter password for account: ")
		accountPassword, _ := reader.ReadString('\n')
		accountPassword = strings.Trim(accountPassword, "\n")

		gsSrv.Add(accountName, accountPassword)

	case "delete":
		fmt.Println("Delete mode...")
		fmt.Print("Enter account name or site URL: ")
		accountName, _ := reader.ReadString('\n')
		accountName = strings.Trim(accountName, "\n")
		err := gsSrv.Delete(accountName)
		if err != nil {
			fmt.Println("FAILED")
		}
	case "search":
		fmt.Println("Search mode...")
		fmt.Print("Enter account name or site URL: ")
		accountName, _ := reader.ReadString('\n')
		accountName = strings.Trim(accountName, "\n")
		accounts, err := gsSrv.Search(accountName)
		if err != nil {
			fmt.Println("FAILED")
		}
		fmt.Println(accounts)
	default:
		fmt.Println("Unrecognized command...")
	}
}
