package explorer

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func fetchBlock(client *ethclient.Client, arg string) (*types.Block, error) {
	ctx := context.Background()
	if arg == "latest" {
		block, err := client.BlockByNumber(ctx, nil)
		if err != nil {

			return nil, err
		}
		return block, nil
	}

	num, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("Wrong number block, couldn't be parsed into int64.")
	}
	block, err := client.BlockByNumber(ctx, big.NewInt(num))
	return block, nil
}

func PrintBlock(block *types.Block) {
	fmt.Println("--------------------------------------------------")
	fmt.Printf("Block Number: %v\n", block.Number())
	fmt.Printf("Hash: %s\n", block.Hash())
	fmt.Printf("Parent Hash: %s\n", block.ParentHash())
	fmt.Printf("Miner: %s\n", block.Coinbase())
	fmt.Printf("Gas Used: %d\n", block.GasUsed())
	fmt.Printf("Transactions: %d\n", len(block.Transactions()))
	fmt.Println("--------------------------------------------------")
}
