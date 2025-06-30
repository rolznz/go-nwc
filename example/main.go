package main

import (
	"context"
	"fmt"
	"os"

	"github.com/rolznz/go-nwc/nip47"
)

func main() {
	fmt.Println("Hello, World!")

	nwcUri := os.Getenv("NWC_URI")
	if nwcUri == "" {
		fmt.Printf("Please set a NWC_URI and try again")
		return
	}

	client, err := nip47.NewNWCClientFromURI(context.TODO(), nwcUri, nil)
	if err != nil {
		fmt.Printf("Failed to create NWC client: %v", err)
		return
	}

	// Get wallet info
	info, err := client.GetInfo(context.TODO())
	if err != nil {
		fmt.Printf("failed to get info: %v", err)
		return
	}
	fmt.Println("Wallet Alias:", info.Alias)

	// Check balance
	balance, _ := client.GetBalance(context.TODO())
	fmt.Println("Balance:", balance.Balance, "msat")
}
