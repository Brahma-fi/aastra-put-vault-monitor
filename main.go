package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"github.com/pradeep-selva/aastra-put-vault-monitor/migrations"
	"github.com/pradeep-selva/aastra-put-vault-monitor/utils"
)

const TIME_FOR_ONE_BLOCK = 13
const TIME_RANGE = 2 * 60 * 60

var db *sql.DB

type BlockRange struct {
	start *big.Int
	end   *big.Int
}

func readAndProcessLogs(
	client *ethclient.Client,
	envVarName string,
	event string,
	blockRange BlockRange,
	eventName string,
) {
	contractAddress := common.HexToAddress(utils.GetEnvVar(envVarName))
	eventsQuery := ethereum.FilterQuery{
		FromBlock: blockRange.start,
		ToBlock:   blockRange.end,
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), eventsQuery)
	utils.CheckError(err, "Could not retrieve vault logs")

	utils.LogInfo(fmt.Sprintf("Loaded contract and retrieved %s logs", eventName))

	for _, vLog := range logs {
		if vLog.Topics[0].Hex() == event {
			block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(vLog.BlockNumber)))
			utils.CheckError(err, "Could not read block")

			INSERTION_COMMAND := fmt.Sprintf(
				"INSERT INTO brahmafi.ethput_activity values (FROM_UNIXTIME(%d), %d, \"%s\", \"%s\")",
				block.Time(),
				block.Number(),
				eventName,
				vLog.TxHash.Hex(),
			)
			log.Println(INSERTION_COMMAND)
			_, err = db.Exec(INSERTION_COMMAND)
			if err != nil {
				log.Println("ERROR: Could not insert into db: " + vLog.TxHash.Hex())
			}

			utils.LogInfo("Successfully inserted into db: " + vLog.TxHash.Hex())
		}
	}
}

func main() {
	// load .env
	err := godotenv.Load(".env")
	utils.CheckError(err, "Could not load .env")

	if utils.GetEnvVar("MIGRATE") == "true" {
		migrations.CreateTable()
	}

	// connect to brahmafi db
	db, err = sql.Open("mysql", utils.GetDbURI())
	utils.CheckError(err, "Could not connect to db")

	utils.LogInfo("Connected to db: brahmafi")

	// connect to infura endpoint: mainnet
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
		"ROUTER_ADDRESS",
		utils.Events.Rebalance,
		BlockRange{olderBlock, currentBlock},
		"rebalance",
	)
	readAndProcessLogs(
		client,
		"AASTRA_VAULT_ADDRESS",
		utils.Events.CompoundFee,
		BlockRange{olderBlock, currentBlock},
		"compoundFee",
	)
}
