package main

import (
	"flag"
	"go-eth-explorer/explorer"
	"log/slog"
	"os"
)

func main() {
	cmd := flag.String("cmd", "block", "command: block or txs")
	arg := flag.String("arg", "latest", "argument for command: latest or block number")
	flag.Parse()
	client, err := explorer.NewClient("")
	if err != nil {
		slog.Error("Error when creating new eth client:", "error", err)
		os.Exit(1)
	}
}
