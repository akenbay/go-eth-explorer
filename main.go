package main

import "flag"

func main() {
	cmd := flag.String("cmd", "block", "command: block or txs")
	arg := flag.String("arg", "latest", "argument for command: latest or block number")
	flag.Parse()

}
