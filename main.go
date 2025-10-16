package main

import (
	"flag"
	"fmt"
	"go-eth-explorer/explorer"
	"log/slog"
	"os"
)

func main() {
	cmd := flag.String("cmd", "block", "command: block or txs")
	arg := flag.String("arg", "latest", "argument for command: latest or block number")
	flag.Parse()

	client, err := explorer.NewClient("https://cloudflare-eth.com")
	if err != nil {
		slog.Error("Error when creating new eth client:", "error", err)
		os.Exit(1)
	}
	defer client.Close()

	block, err := explorer.FetchBlock(client, *arg)
	if err != nil {
		slog.Error("Error when fetching block:", "error", err)
		os.Exit(1)
	}

	switch *cmd {
	case "block":
		explorer.PrintBlock(block)
	case "txs":
		count := explorer.PrintLatestTxs(block)
		fmt.Printf("Displayed %d transactions.\n", count)
	default:
		fmt.Println("Unknown command. Use -cmd=block or -cmd=txs")
		os.Exit(1)
	}

}
