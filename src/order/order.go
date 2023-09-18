package order

import (
	"context"
	"fmt"
	"strings"
	"tradingBot/src/account"
	"tradingBot/src/logging"
	"tradingBot/src/tools"

	"github.com/adshao/go-binance/v2"
)

func BuyMarketOrder(client *binance.Client, symbol string) (*binance.CreateOrderResponse, error) {
    amount, _ := account.GetBalance(client, "USDT")
    quoteOrderQty := fmt.Sprintf("%v", amount)

    order, err := client.NewCreateOrderService().Symbol(symbol).
        Side(binance.SideTypeBuy).
        Type(binance.OrderTypeMarket).
        QuoteOrderQty(quoteOrderQty).
        Do(context.Background())

    if err != nil {
        return nil, err
    } else {
        return order, nil
    }
}

func SellMarketOrder(client *binance.Client, symbol string) (*binance.CreateOrderResponse, error) {
    target := strings.Replace(symbol, "USDT", "", 1)
    amount, _ := account.GetBalance(client, target)

    step := account.GetStepSymbol(client, symbol)
    quantity := tools.RoundDown(amount, step)
    quantityString := fmt.Sprintf("%v", quantity)

    order, err := client.NewCreateOrderService().Symbol(symbol).
        Side(binance.SideTypeSell).
        Type(binance.OrderTypeMarket).
        Quantity(quantityString).
        Do(context.Background())

    if err != nil {
        return nil, err
    } else {
        return order, nil
    }
}
func SellOcoOrder(client *binance.Client, symbol string, price, stopPrice, stopLimitPrice float64) (*binance.CreateOCOResponse, error) {
    target := strings.Replace(symbol, "USDT", "", 1)
    amount, _ := account.GetBalance(client, target)

    step := account.GetStepSymbol(client, symbol)
    stepUSDT := account.GetStepUSDTSymbol(client, symbol)
    quantity := tools.RoundDown(amount, step)
    quantityString := fmt.Sprintf("%v", quantity)

    priceString := fmt.Sprintf("%v", tools.RoundDown(price, stepUSDT))
    stopString := fmt.Sprintf("%v", tools.RoundDown(stopPrice, stepUSDT))
    limitString := fmt.Sprintf("%v", tools.RoundDown(stopLimitPrice, stepUSDT))

    logging.LogError("[DEBUG] %f", price)
    logging.LogError("[DEBUG] %f", stopPrice)
    logging.LogError("[DEBUG] %f", stopLimitPrice)

    order, err := client.NewCreateOCOService().
        Symbol(symbol).
		Side(binance.SideTypeSell).
		Quantity(quantityString).
		Price(priceString).
		StopPrice(stopString).
		StopLimitPrice(limitString).
		Do(context.Background())

    logging.LogError("[DEBUG] %v", order)

    if err != nil {
        return nil, err
    } else {
        return order, nil
    }
}



