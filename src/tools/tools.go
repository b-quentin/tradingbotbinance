package tools

import (
	"fmt"
	"math"
	"strconv"

	"github.com/adshao/go-binance/v2"
)

func KlineBySymbolAndTime(kline *binance.WsKline, symbol string, interval string) (bool, *binance.WsKline){
    if(kline.Symbol == symbol && kline.Interval == interval) {
        return true , kline
    } else {
        return false, nil
    }
}

func StringToFloat(data string) float64 {
    payback, err := strconv.ParseFloat(data, 64)

    if err != nil {
        fmt.Println(err)
    }

    return payback
}

// Rounds like 12.3496 -> 12.34
func RoundDown(val float64, precision int) float64 {
    return math.Floor(val*(math.Pow10(precision))) / math.Pow10(precision)
}

