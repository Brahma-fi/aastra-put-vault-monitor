## Aastra PUT Vault Monitor

A bot to monitor and persistently track compounding and rebalance events on Aastra ETH-PUT vault.

## To Run

1. `go mod vendor`
2. `go run main.go`

(or)

1. `docker build -t <name>:1.0 .`
2.

```
docker run -e AASTRA_VAULT_ADDRESS='<vault address>' \
-e ROUTER_ADDRESS='<router address>' \
-e INFURA_ENDPOINT='<infura endpoint>' \
-e MIGRATE='<migrate>' \
-e DB_USERNAME='<username>' \
-e DB_PASSWORD='<password>'
-e DB_ENDPOINT='<endpoint>'
-e DB_PORT='<port>'
-e DB_NAME='<db name>'
```

- Set `MIGRATION` to true to create DBs before running script.
