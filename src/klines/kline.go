package klines

import (
	"fmt"
	"time"

	//"math"
	"strconv"

	"github.com/adshao/go-binance/v2"
//	"github.com/b-quentin/indicator"
)

type Kline struct {
    OpenTime time.Time
    ClosedTime time.Time
    OpenPrice float64
    ClosedPrice float64
    LowPrice float64
    HighPrice float64
    Rsi Rsi
    StockRsi StockRsi
    Ema1 float64
    Ema2 float64
    Macd float64
    Signal float64
    Histo float64
    MiddleBand float64
    UpperBand float64
    LowerBand float64
    AtrLower float64
    Adx float64
    Plus float64
    Minus float64
    ConversionBase float64
    BaseLigne float64
    LeadingSpanA float64
    LeadingSpanB float64
    LaggingSpan float64
}

func (kl *Kline) New(k *binance.WsKline) {
    kl.OpenTime = timespanToDate(k.StartTime)
    kl.ClosedTime = timespanToDate(k.EndTime)
    kl.OpenPrice = stringToFloat(k.Open)
    kl.ClosedPrice = stringToFloat(k.Close)
    kl.HighPrice = stringToFloat(k.High)
    kl.LowPrice = stringToFloat(k.Low)
}

func (kl *Kline) ShowKline() {
    fmt.Println("OpenTime:", kl.OpenTime, 
        "Open:", kl.OpenPrice,
        "Close:", kl.ClosedPrice,
        "High:", kl.HighPrice,
        "Low:", kl.HighPrice,
        "CloseTime:", kl.ClosedTime,
        "Ema1", kl.Ema1,
        "Ema2", kl.Ema2,
        "LowerBand:", kl.LowerBand,
        "ATR Lower:", kl.AtrLower,
        "Rsi:", kl.Rsi.getRsi(),
        "K:", kl.StockRsi.GetK(),
        "D:", kl.StockRsi.GetD(),
        "ADX:", kl.Adx,
        "Plus:", kl.Plus,
        "Minus", kl.Minus,
        "Macd:", kl.Macd,
        "Signal:", kl.Signal,
        "diff:", kl.Macd - kl.Signal,
        "LeadingSpanA:", kl.LeadingSpanA,
        "LeadingSpanB:", kl.LeadingSpanB,
        )
}

func stringToFloat(data string) float64 {
    payback, err := strconv.ParseFloat(data, 64)

    if err != nil {
        fmt.Println(err)
    }

    return payback
}

func (kl *Kline) GetKline() *Kline {
    return kl
}


