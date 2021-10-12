package migrations

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/pradeep-selva/aastra-put-vault-monitor/utils"
)

func CreateTable() {
	err := godotenv.Load(".env")
	utils.CheckError(err, "Could not load .env")
	
	db, err := sql.Open("mysql", utils.GetDbURI())
	utils.CheckError(err, "Could not connect to db")
	
	ETHPUT_ACTIVITY_TABLE_COMMAND := "create table ethput_activity(timeStamp datetime, blockNumber float, activityType varchar(255), txnHash varchar(255))"
	_,err = db.Exec(ETHPUT_ACTIVITY_TABLE_COMMAND)
	utils.CheckError(err, "Could not create table: ethput_activity")

	utils.LogInfo("Created table: brahmafi.ethput_activity")
}