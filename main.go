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

func main(){
	err := godotenv.Load(".env")
	utils.CheckError(err, "Could not load .env")

	client, err := ethclient.Dial(utils.GetEnvVar("INFURA_ENDPOINT"))
	utils.CheckError(err, "Could not initialize eth client")

	utils.LogInfo("Connected to infura node")
	
	contractAddress := common.HexToAddress(utils.GetEnvVar("AASTRA_VAULT_ADDRESS"))
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(13371034),
		ToBlock: big.NewInt(13396011),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	utils.CheckError(err, "Could not retrieve logs")

	utils.LogInfo("Loaded contract and retrieved logs")

	_, err = abi.JSON(strings.NewReader(string(AastraVault.AastraVaultABI)))
	utils.CheckError(err, "Could not read contract ABI")

	for _, vLog := range logs {
		log.Println(vLog.Topics[0].Hex())
	}
}