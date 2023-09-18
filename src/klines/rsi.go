package klines

import (
	"math"

	"github.com/b-quentin/indicator"
)

type Rsi struct {
    Value float64
}

func (kls *Klines) SetKlinesRSI(period int) {
    _, rsi := indicator.RsiPeriod(period, kls.Closed)
    
    for i := range kls.Kline {
        accessKline := &kls.Kline[i]
        accessKline.SetRsi(rsi[i])    
    }
}

func (kls *Klines) SetKlinesStockRSI() {
    stock := indicator.StockRsi{}
    stock.RsiPeriod = 14
    stock.StockPeriod = 14
    stock.KPeriod = 3
    stock.DPeriod = 3

    k, d := stock.StochasticOscillatorRSI(kls.Closed)
    
    startAt := len(kls.Closed) - len(k)
    for i := startAt; i < len(kls.Closed); i++ {
        step := i - startAt
        accessKline := &kls.Kline[i]
        accessKline.SetStockRsi(k[step],d[step])
    }
}

func (kl *Kline ) SetRsi(rsi float64) {
    kl.Rsi.SetRsi(rsi)
}

func (kl *Kline) SetStockRsi(k, d float64) {
    kl.StockRsi.SetStockRsi(k, d)
}

func (rsi *Rsi) roundDown(precision int) {
    rsi.Value = math.Floor(rsi.Value*(math.Pow10(precision))) / math.Pow10(precision)
}

func (rsi *Rsi ) SetRsi(data float64) {
    rsi.Value = data
    rsi.roundDown(2)
}

func (rsi *Rsi) getRsi() float64 {
    return rsi.Value
}

func (rsi *Rsi) isOverbought(limit float64) bool {
    if rsi.Value >= limit {
        return true
    } else {
        return false
    }
}

func (rsi *Rsi) isOversold(limit float64) bool {
    if rsi.Value <= limit {
        return true
    } else {
        return false
    }
}
