# go-eth-explorer

A lightweight block explorer in Ethereum written in Go.
Designed to learn and demonstrate how to interact with the net with go-ethereum SDK.

It connects to any Ethereum RPC endpoint and gets information about specific or latest block (block number, hash, parent hash, miner, gas used and number of transactions). It also can inspect latest 10 transactions of specific or latest block.

## 🧱 Project Structure

```

go-eth-explorer/
├── main.go
├── explorer/
│   ├── client.go
│   ├── block.go
│   ├── tx.go
│   └── utils.go
├── go.mod
├── .env.example
└── README.md

```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/akenbay/go-eth-explorer.git
   cd go-eth-explorer
   ```
2. Initialize modules:
   ```bash
   go mod tidy
   ```
3. Create a .env file to store your RPC URL

## 🧩 Usage

### Fetch the latest block

```bash
go run main.go -cmd=block -arg=latest
```

### Fetch a specific block

```bash
go run main.go -cmd=block -arg=21000000
```

### List transactions in a block

```bash
go run main.go -cmd=txs -arg=latest
```

Example output:

```
--------------------------------------------------
Block Number: 21312487
Hash: 0x89a...
Parent Hash: 0x2af...
Miner: 0xabc...f21
Gas Used: 18124000
Transactions: 189
--------------------------------------------------
[1] 0xa12...
[2] 0xb45...
Displayed 10 transactions.
```

---

## 🔐 Environment Variables

| Variable  | Description                                   |
| --------- | --------------------------------------------- |
| `ETH_API` | Ethereum RPC endpoint (Infura, Alchemy, etc.) |

Example `.env`:

```bash
ETH_API=https://mainnet.infura.io/v3/YOUR_PROJECT_ID
```

## 🧑‍💻 Author

**Aibar Kenbay**
[GitHub](https://github.com/akenbay)
