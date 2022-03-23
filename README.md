# CryptoMarket-go
[main page](https://www.cryptomkt.com/)


[sign up in CryptoMarket](https://www.cryptomkt.com/account/register).

# Installation
To install the Cryptomarket client
```
 go get github.com/cryptomarket/cryptomarket-go
```
# Documentation
This sdk makes use of the [api version 2](https://api.exchange.cryptomkt.com/v2) of cryptomarket


# Quick Start

## rest client
```go
import (
	"context"

    "github.com/cryptomarket/cryptomarket-go/args"
    "github.com/cryptomarket/cryptomarket-go/rest"
)

// instance a client
let apiKey="AB32B3201"
let apiSecret="21b12401"
client := rest.NewClient(apiKey, apiSecret)
ctx := context.Background()

// get currencies
currencies, err := client.GetCurrencies(ctx)

// get order books
orderBook, err := client.GetOrderBook(ctx, args.Symbol("EOSETH"))

// get your wallet balances
accountBalanceList, err := client.GetWalletBalances(ctx)

// get your trading balances
tradingBalanceList, err := client.GetSpotTradingBalances(ctx)

// move balance from wallet to spot trading
result, err := client.TransferMoneyFromAccountBalanceToTradingBalance(
  ctx,
  args.Currency("ETH"),
  args.Amount("3.2"),
  args.Source(args.AccountWallet),
  args.Destination(args.AccountSpot),
)

// get your active orders
ordersList, _ := client.GetAllActiveSpotOrders(ctx, args.Symbol("EOSETH"))

// create a new order
order, err := client.CreateSpotOrder(ctx, args.Symbol("EOSETH"), args.Side(args.SideTypeBuy), args.Quantity("10"), args.Price("10"))
```

## websocket client
*work in progress*
## arguments and constants of interest
all the arguments for the clients are in the args package, as well as the custom types for the arguments. check the package documentation, and the method documentation of the clients for more info.

# Checkout our other SDKs
<!-- agregar links -->
python sdk

nodejs sdk

java sdk

ruby sdk