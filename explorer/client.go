package explorer

import (
	"log/slog"

	"github.com/ethereum/go-ethereum/ethclient"
)

func newClient(url string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		slog.Error("Error when creating new eth client:", "error", err)
		return nil, err
	}
	return client, nil
}
