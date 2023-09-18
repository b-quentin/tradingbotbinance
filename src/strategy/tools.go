package strategy

import (
	"sort"
	"sync"
	"tradingBot/src/klines"

	"github.com/adshao/go-binance/v2"
)

func CheckEma(client *binance.Client,SYMBOL string, INTERVAL string) (bool, string) {
    kl := klines.Klines{}
    kl.SetKlines(client, SYMBOL, INTERVAL, 500)
    kl.SetEma1(50)
    kl.SetEma2(200)

    lastKl := kl.GetLastKline()

    if (lastKl.IsEma1AboveEma2()) {
        return true, SYMBOL
    } else {
        return false, ""
    }
}

func GetListSymbols(client *binance.Client, symbols []string, interval string) []string {
	// Créer un groupe de wait pour attendre la fin de toutes les goroutines
	var wg sync.WaitGroup
    nbr := len(symbols)

    var SymbolsChecked []string
    wg.Add(nbr)
    for _, sy := range symbols {
        go func(symbol string) {
            defer wg.Done()
            emaBool, sbl := CheckEma(client, symbol, interval)
            if emaBool {
                SymbolsChecked = append(SymbolsChecked, sbl)
            }
        }(sy)
    }

	wg.Wait()
    
    sort.Strings(SymbolsChecked)
    return SymbolsChecked
}
