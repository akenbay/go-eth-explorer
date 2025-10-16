package explorer

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

func PrintLatestTxs(block *types.Block) int {
	count := len(block.Transactions())

	fmt.Printf("Transactions in block %v:\n", block.Number())
	fmt.Println("--------------------------------------------------")
	for i, tx := range block.Transactions() {
		if i >= 10 {
			break // print only first 10
		}
		fmt.Printf("[%d] %s\n", i+1, tx.Hash().Hex())
	}
	fmt.Println("--------------------------------------------------")
	return count
}
