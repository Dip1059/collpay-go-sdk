# Collpay Go Sdk
This is the official golang sdk for Collpay Payment Gateway

[![GoDoc](https://godoc.org/github.com/Dip1059/collpay-go-sdk?status.svg)](https://godoc.org/github.com/Dip1059/collpay-go-sdk)

# Installation
go get github.com/Dip1059/collpay-go-sdk

``` go
import (
    "github.com/Dip1059/collpay-go-sdk"
)
```

# Usage

```go

//Configure environment
    err := collpay.ConfigureEnv(&collpay-go-sdk.Config{
            PublicKey: "xxxxxxxxxxx", //Merchant's public api key
            Env: collpay-go-sdk.ENV_SANDBOX, //or ENV_PRODUCTION
        })

//Get exchange rate
    exchange, err := collpay_go_sdk.GetExchangeRate("USD", "BTC")
    fmt.Println(exchange.Rate)

//Create Transaction
    tr, err := collpay_go_sdk.CreateTransaction(&collpay_go_sdk.Transaction{
            PaymentCurrency:"BTC",
            OrderCurrency:"USD",
            OrderAmount:9.56,
            PayerName:"xxxx",
            PayerEmail:"xxxx",
            PayerPhone: "xxxx",
            PayerAddress: "xxxx",
            IpnUrl:"xxxx",
            IpnSecret:"xxxx", //Any random secret string of your's, It can be empty.
            SuccessUrl:"xxxx",
            CancelUrl:"xxxx",
            Cart:`{"item_name":"t-shirt","item_number":"23","invoice":"SDF-453672-PMT"}`, //Send any data format like json
            WebhookData: `{"order_id":"ABC12345-12"}`, //Send any data format like json
        })
    fmt.Println(tr.ID)

//Get Transaction by ID
    trx, err := collpay_go_sdk.GetTransaction("xxxx")
    fmt.Println(trx.ID)

```