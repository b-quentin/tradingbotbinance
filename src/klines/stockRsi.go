package klines

import "math"

type StockRsi struct {
    K float64
    D float64
}

func (sr *StockRsi) roundDown(precision int) {
    sr.K = math.Floor(sr.K*(math.Pow10(precision))) / math.Pow10(precision)
    sr.D = math.Floor(sr.D*(math.Pow10(precision))) / math.Pow10(precision)
}

func (sr *StockRsi) lessThanZero(data float64) float64 {
    if data < 0 {
        return 0
    } else {
        return data
    }
}

func (sr *StockRsi) SetStockRsi(k, d float64) {
    sr.K = sr.lessThanZero(k)
    sr.D = sr.lessThanZero(d)

    sr.roundDown(2)
}

func (sr *StockRsi) GetK() float64 {
    return sr.K
}

func (sr *StockRsi) GetD() float64 {
    return sr.D
}

func (sr *StockRsi) IsUpTrend() bool {
    if sr.K > sr.D {
        return true
    } else {
        return false
    }
}

func (sr *StockRsi) IsDownTrend() bool {
    if sr.K < sr.D {
        return true
    } else {
        return false
    }
}

func (sr *StockRsi) IsBelowLower() bool {
    if sr.K < 20 && sr.D < 20 {
        return true
    } else {
        return false
    }
}

func (sr *StockRsi) IsAboveLower() bool {
    if sr.K > 20 && sr.D > 20 {
        return true
    } else {
        return false
    }
}

func (sr *StockRsi) IsAboveTop() bool {
    if sr.K > 80 && sr.D > 80 {
        return true
    } else {
        return false
    }
}
