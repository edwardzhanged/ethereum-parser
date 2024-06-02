# ethereum-parser

## Design
Overall, it is a restful api implementation of parser.

A multi-threaded monitoring program will be launched to continuously monitor the transactions on the latest blockchain related to the subscribed addresses.

### GetCurrentBlock

Get the latest block from [eth_getBlockByNumber](https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_getblockbynumber])

### Subscribe
Add an address into storage. Will deny duplicated address.

### Get Transactions
Return all transactions releated to an adress after subscribed. Transactions before subscription are NOT  INCLUDED. 


## How to run

### Requirements

Golang version == 1.22.1

### Build 
```bash
make build 
```

### Run
```bash
./ethereum-parser 
```

### Test
```bash
make test
```


## Usages

To get current block
```bash
curl -X GET 'http://localhost:8080/currentBlock'
```

To subscribe an address
```bash
curl -X POST -d "address=0x1111e3ef0b6ae32e14a55e0e7cd9b8505177c2bf" http://localhost:8080/subscribe
```

To get transactions from an address
```bash
curl -X GET "http://localhost:8080/getTransactions?address=0x1111e3ef0b6ae32e14a55e0e7cd9b8505177c2bf"
```