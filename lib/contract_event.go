package lib

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func GetAllEvents(fromBlock int64, logHandler func(log types.Log), endHandler func(lastLog *types.Log)) {
	config, err := NewConfig()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	client, err := rpc.Dial(config.Contract.HTTP_RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	ethClient := ethclient.NewClient(client)
	contractAddress := common.HexToAddress(config.Contract.CONTRACT_ADDRESS)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		FromBlock: big.NewInt(fromBlock),
	}

	logs, err := ethClient.FilterLogs(context.Background(), query)

	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

	for _, vLog := range logs {
		logHandler(vLog)
	}

	if len(logs) > 0 {
		endHandler(&logs[len(logs)-1])
	} else {
		endHandler(nil)
	}
}

func GetLastetBlock() (*types.Header, error) {
	config, err := NewConfig()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	client, err := rpc.Dial(config.Contract.HTTP_RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	ethClient := ethclient.NewClient(client)

	block, err := ethClient.HeaderByNumber(context.Background(), nil)

	if err != nil {
		log.Fatalf("Failed to filter logs: %v", err)
	}

	return block, nil
}

func ListenToNewBlocks(logHandler func(log types.Log)) {
	config, err := NewConfig()
	if err != nil {
		log.Fatalf("Failed to get config: %v", err)
	}

	client, err := rpc.Dial(config.Contract.WS_RPC_URL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	ethClient := ethclient.NewClient(client)
	contractAddress := common.HexToAddress(config.Contract.CONTRACT_ADDRESS)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to new logs: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Error in new log subscription: %v", err)
		case vLog := <-logs:
			logHandler(vLog)
		}
	}
}
