package main

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	AastraVault "github.com/pradeep-selva/aastra-put-vault-monitor/generated/aastra-vault"
	"github.com/pradeep-selva/aastra-put-vault-monitor/utils"
)

const TIME_FOR_ONE_BLOCK = 13
const TIME_RANGE = 2 * 60 * 60

type eventProcessor func()
type BlockRange struct {
	start *big.Int
	end   *big.Int
}

func readAndProcessLogs(
	client *ethclient.Client,
	envVarName string,
	event string,
	blockRange BlockRange,
	processEvent eventProcessor,
) {
	contractAddress := common.HexToAddress(utils.GetEnvVar(envVarName))
	eventsQuery := ethereum.FilterQuery{
		FromBlock: blockRange.start,
		ToBlock:   blockRange.end,
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), eventsQuery)
	utils.CheckError(err, "Could not retrieve vault logs")

	utils.LogInfo("Loaded contract and retrieved logs")

	_, err = abi.JSON(strings.NewReader(string(AastraVault.AastraVaultABI)))
	utils.CheckError(err, "Could not read contract ABI")

	for _, vLog := range logs {
		if vLog.Topics[0].Hex() == event {
			processEvent()
		}
	}
}

func main() {
	err := godotenv.Load(".env")
	utils.CheckError(err, "Could not load .env")

	client, err := ethclient.Dial(utils.GetEnvVar("INFURA_ENDPOINT"))
	utils.CheckError(err, "Could not initialize eth client")

	utils.LogInfo("Connected to infura node")

	header, err := client.HeaderByNumber(context.Background(), nil)
	utils.CheckError(err, "Could not retrieve headers")

	currentBlock := header.Number
	// block from 2 hours ago
	olderBlock := big.NewInt(0).Sub(currentBlock, big.NewInt(TIME_RANGE/TIME_FOR_ONE_BLOCK))

	readAndProcessLogs(
		client,
		"AASTRA_VAULT_ADDRESS",
		utils.Events.CompoundFee,
		BlockRange{olderBlock, currentBlock},
		func() { log.Println("CompoundFee") },
	)
	readAndProcessLogs(
		client,
		"ROUTER_ADDRESS",
		utils.Events.Rebalance,
		BlockRange{olderBlock, currentBlock},
		func() { log.Println("Rebalance") },
	)
}
