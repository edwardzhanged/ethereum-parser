# ethereum-parser

## Design
A restful api implementation of parser.


## How to run

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
curl --location 'http://localhost:8080/currentBlock'
```

To subscribe an address
```bash
curl --location 'http://localhost:8080/subscribe?address=0x1111e3ef0b6ae32e14a55e0e7cd9b8505177c2bf'
```

To get transactions from an address
```bash
curl --location 'http://localhost:8080/getTransactions?address=0x1111e3ef0b6ae32e14a55e0e7cd9b8505177c2bf'
```