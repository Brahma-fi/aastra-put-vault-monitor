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
	
	vaultAddress := common.HexToAddress(utils.GetEnvVar("AASTRA_VAULT_ADDRESS"))
	vaultEventsQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(13371034),
		ToBlock: big.NewInt(13396011),
		Addresses: []common.Address{vaultAddress},
	}

	vaultLogs, err := client.FilterLogs(context.Background(), vaultEventsQuery)
	utils.CheckError(err, "Could not retrieve vault logs")

	utils.LogInfo("Loaded Vault contract and retrieved logs")

	_, err = abi.JSON(strings.NewReader(string(AastraVault.AastraVaultABI)))
	utils.CheckError(err, "Could not read contract ABI")

	for _, vLog := range vaultLogs {
		if vLog.Topics[0].Hex() == utils.Events.CompoundFee {
			log.Println("CompoundFee")
		}
	}

	routerAddress := common.HexToAddress(utils.GetEnvVar("ROUTER_ADDRESS"))
	routerEventsQuery := ethereum.FilterQuery{
		FromBlock: big.NewInt(13371034),
		ToBlock: big.NewInt(13396011),
		Addresses: []common.Address{routerAddress},
	}

	routerLogs, err := client.FilterLogs(context.Background(), routerEventsQuery)
	utils.CheckError(err, "Could not retrieve router logs")

	utils.LogInfo("Loaded Router contract and retrieved logs")

	for _, vLog := range routerLogs {
		if vLog.Topics[0].Hex() == utils.Events.Rebalance {
			log.Println("Rebalance")
		}
	}
}