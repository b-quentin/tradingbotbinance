package strategy

// Porte d'entré de tte les stratégies 
// CheckEMA pour ttes les heures
// Et séléctionne les SYMBOLS les plus BULL

import (
	"fmt"
	"sync"
	"tradingBot/src/account"
	"tradingBot/src/tools"
	"tradingBot/src/logging"

	"github.com/adshao/go-binance/v2"
)

func Launch(client *binance.Client) {
    symbols := account.GetSymbols()

    myIsTrade := &InTrade{InTrade: false, Symbol: "", TakeProfit: 0, StopLoss: 0, ValueAtBuy: 0}
    myLog := &logging.Logging{}
    stopChan := make(chan struct{})

    listSymbolDaily := GetListSymbols(client, symbols, "1d")
    listSymbol4h := GetListSymbols(client, listSymbolDaily, "4h")
    listSymbol1h := GetListSymbols(client, listSymbol4h, "1h")
    listSymbol15m := GetListSymbols(client, listSymbol1h, "15m")
    listSymbol5m := GetListSymbols(client, listSymbol15m, "5m")


    myLog.LogData("[ALL][HOUR] %s", listSymbol5m)

    var wg sync.WaitGroup

    nbr := len(listSymbol5m)

    wg.Add(nbr)
    for _, sy := range listSymbol5m {
        go func(symbol string, isTrade I_InTrade, isLog logging.I_Logging) {
            st := Strategy4{}
            select {
            case <-stopChan:
                return // Terminer la goroutine si le canal a été fermé
            default:
                st.Launch(client, symbol, "5m", myIsTrade, isLog)
            }
        }(sy, myIsTrade, myLog)
    }

    wsKlineHandler := func(event *binance.WsKlineEvent) {
        boolDaily, klineDaily := tools.KlineBySymbolAndTime(&event.Kline ,"BTCUSDT", "1d")
        bool4h, kline4h := tools.KlineBySymbolAndTime(&event.Kline ,"BTCUSDT", "4h")
        bool1h, kline1h := tools.KlineBySymbolAndTime(&event.Kline ,"BTCUSDT", "1h")
        bool15m, kline15m := tools.KlineBySymbolAndTime(&event.Kline ,"BTCUSDT", "15m")
        bool5m, kline5m := tools.KlineBySymbolAndTime(&event.Kline ,"BTCUSDT", "5m")


        if boolDaily && klineDaily.IsFinal {
            listSymbolDaily = GetListSymbols(client, symbols, "1d")
        }

        if bool4h && kline4h.IsFinal {
            listSymbol4h = GetListSymbols(client, listSymbolDaily, "4h")
        }

        if bool1h && kline1h.IsFinal {
            listSymbol1h = GetListSymbols(client, listSymbol4h, "1h")

            close(stopChan)

            myLog.LogData("[ALL][HOUR] %s", listSymbol5m)

            nbr = len(listSymbol5m)
            wg.Add(nbr)
            for _, sy := range listSymbol5m {
                go func(symbol string, isTrade I_InTrade, isLog logging.I_Logging) {
                    st := Strategy4{}
                    select {
                        case <-stopChan:
                        return // Terminer la goroutine si le canal a été fermé
                    default:
                        st.Launch(client, symbol, "5m", myIsTrade, isLog)
                    }
                }(sy, myIsTrade, myLog)
            }
        }

        if bool15m && kline15m.IsFinal {
            listSymbol15m = GetListSymbols(client, listSymbol1h, "15m")
        }

        if bool5m && kline5m.IsFinal {
            data := fmt.Sprintf("[ALL][INTRADE] %t, Symbol: %s", myIsTrade.GetInTrade(), myIsTrade.GetSymbol())
            myLog.LogData(data)
            listSymbol5m = GetListSymbols(client, listSymbol5m, "5m")
        }
    }

    errHandler := func(err error) {
        fmt.Println(err)
    }

    // Pour avoir un trigger par horaire
    symbolDaily := map[string]string{"BTCUSDT": "1d"}
    symbol4h := map[string]string{"BTCUSDT": "4h"}
    symbol1h := map[string]string{"BTCUSDT": "1h"}
    symbol15m := map[string]string{"BTCUSDT": "15m"}
    symbol5m := map[string]string{"BTCUSDT": "5m"}

    doneSymbolDaily, _, _ := binance.WsCombinedKlineServe(symbolDaily, wsKlineHandler, errHandler)
    doneSymbol4h, _, _ := binance.WsCombinedKlineServe(symbol4h, wsKlineHandler, errHandler)
    doneSymbol1h, _, _ := binance.WsCombinedKlineServe(symbol1h, wsKlineHandler, errHandler)
    doneSymbol15m, _, _ := binance.WsCombinedKlineServe(symbol15m, wsKlineHandler, errHandler)
    doneSymbol5m, _, _ := binance.WsCombinedKlineServe(symbol5m, wsKlineHandler, errHandler)

    <-doneSymbolDaily
    <-doneSymbol4h
    <-doneSymbol1h
    <-doneSymbol15m
    <-doneSymbol5m
}
