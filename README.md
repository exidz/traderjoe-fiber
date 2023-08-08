
# Traderjoe Price pool feed Rest API

Traderjoe price pool feed Rest API, developed in Golang using Fiber framework. Support v2.1 Liquidity Book Contracts, v2 Liquidity Book Contracts, and v1 Trader Joe Contracts across three major chains: Avalanche, Arbitrum, and Binance Smart Chain. 




## Swagger Documentation

[Documentation](https://traderjoe-fiber-staging.up.railway.app/swagger/)


## Sequence Diagram

![](https://i.imgur.com/qu5SMf3.jpg)
> Getting a v2.1 Pool Price using v2.1 LB contracts

![](https://i.imgur.com/YBKtSZw.jpg)
> Getting a v2 Pool Price using v2 LB contracts

![](https://i.imgur.com/l96KdZf.jpeg)
> Getting a v1 Pool Price using v1 Joe contracts



## Benchmark

#### Benchmark using go-wrk for single pool price
```bash
Running 10s test @ https://traderjoe-fiber-staging.up.railway.app/avax/v2_1/prices/0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e/0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7/20
  100 goroutine(s) running concurrently
5389 requests in 10.091529806s, 2.37MB read
Requests/sec:		534.01
Transfer/sec:		240.34KB
Avg Req Time:		187.26164ms
Fastest Request:	157.158415ms
Slowest Request:	2.091516792s
Number of Errors:	0

```
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`APP_HOST`

`APP_PORT`

`APP_READ_TIMEOUT`

`AVAX_RPC`

`ARB_RPC`

`BSC_RPC`


## Installation

Run the project using Docker

```bash
  git clone https://github.com/exidz/traderjoe-fiber.git
  cd traderjoe-fiber
  docker-compose up --build
```
Alternatively, you can run locally

```bash
  git clone https://github.com/exidz/traderjoe-fiber.git
  cd traderjoe-fiber
  go mod download
  go build -o traderjoe main.go
  ./traderjoe
```
## Running test

To run tests, run the following command

```bash
  go test -v ./...
```



## API Reference

#### Get v2.1/v2 single pool price (1:1)

```http
  GET /${chain}/${version}/prices/${baseAsset}/${quoteAsset}/${binStep}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `chain` | `string` | **Required**.  chain of the pool (avax, arb or bsc)|
| `version` | `string` | **Required**.  version of the pool (v2 or v2.1)|
| `baseAsset` | `string` | **Required**. ERC20 token contract address|
| `quoteAsset` | `string` | **Required**.  ERC20 token contract address|
| `binStep` | `int` | **Required**.  binStep of the pool|

Sample curl request
```http
curl --location --request GET 'https://traderjoe-fiber-staging.up.railway.app/avax/v2_1/prices/0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7/0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e/20'
```

#### Get v2.1/v2 batch pool price (1:1)

```http
  POST /${chain}/${version}/batch-prices
```
Support 20 batch pool price per request
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `chain` | `string` | **Required**.  chain of the pool (avax, arb or bsc)|
| `version` | `string` | **Required**.  version of the pool (v2 or v2.1)|

| Request body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `data` | `array` | **Required**.  list of baseAsset, quoteAsset and binstep|

Sample curl request

```bash
curl --location --request POST 'https://traderjoe-fiber-staging.up.railway.app/avax/v2_1/batch-prices' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [
        {
            "baseAsset": "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7",
            "quoteAsset": "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e",
            "binStep": 20
        },
        {
            "baseAsset": "0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7",
            "quoteAsset": "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e",
            "binStep": 1
        }
    ]
}'

```

#### Get v1 single pool price (1:1)

```http
  GET /${chain}/v1/prices/${baseAsset}/${quoteAsset}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `chain` | `string` | **Required**.  chain of the pool (avax, arb or bsc)|
| `baseAsset` | `string` | **Required**. ERC20 token contract address|
| `quoteAsset` | `string` | **Required**.  ERC20 token contract address|

Sample curl request
```bash
curl --location --request GET 'https://traderjoe-fiber-staging.up.railway.app/avax/v1/prices/0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7/0xce1bffbd5374dac86a2893119683f4911a2f7814'
```

#### Get v1 batch pool price (1:1)

```bash
  POST /${chain}/v1/batch-prices
```
Support 20 batch pool price per request
| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `chain` | `string` | **Required**.  chain of the pool (avax, arb or bsc)|

| Request body | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `data` | `array` | **Required**.  list of baseAsset, and quoteAsset|

Sample curl request

```bash
curl --location --request POST 'https://traderjoe-fiber-staging.up.railway.app/avax/v1/batch-prices' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [
        {
            "baseAsset": "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7",
            "quoteAsset": "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e"
        },
        {
            "baseAsset": "0x62edc0692bd897d2295872a9ffcac5425011c661",
            "quoteAsset": "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7"
        }
    ]
}'

```




## Lessons Learned

During the project's development with Golang, I discovered the language's exceptional concurrency support, which allowed me to achieve remarkable parallelism and performance through goroutines. However, I encountered a challenge in finding a multicall library for Golang, prompting me to utilize the Trader Joe contract interfaces to efficiently handle multiple contract calls without dedicated multicall support.


## Author

- [@exidz](https://www.github.com/exidz)

