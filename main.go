package main

import (
	"fmt"

	"github.com/OPM123/simple-blockchain/models"
)

func main() {

	bc := models.NewBlockchain()

	bc.AddBlock("Deposit 2 BTC")
	bc.AddBlock("Transfer 1 BTC")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
