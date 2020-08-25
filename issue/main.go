package main

import (
	"fmt"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/alpacahq/alpaca-trade-api-go/stream"
)

func main() {
	os.Setenv(common.EnvApiKeyID, "YOUR_API_KEY_HERE")
	os.Setenv(common.EnvApiSecretKey, "YOUR_API_SECRET_HERE")

	// if err := stream.Register(alpaca.TradeUpdates, tradeHandler); err != nil {
	// 	panic(err)
	// }

	if err := stream.Register(alpaca.AccountUpdates, accountHandler); err != nil {
		panic(err)
	}

	select {}
}

func tradeHandler(msg interface{}) {
	tradeupdate := msg.(alpaca.TradeUpdate)
	fmt.Printf("TRADE UPDATE: %s event received for order %s.\n", tradeupdate.Event, tradeupdate.Order.ID)
}

func accountHandler(msg interface{}) {
	accountupdate := msg.(alpaca.AccountUpdate)
	fmt.Printf("ACCOUNT UPDATE: %s event received for account %s.\n", accountupdate.Event, accountupdate.Account.ID)
}
