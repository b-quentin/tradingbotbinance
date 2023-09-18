package klines

import "github.com/cinar/indicator"

func (kls *Klines) SetBollingerBands() {
    _, _, lowerBand := indicator.BollingerBands(kls.Closed)

    for i := range kls.Kline {
        accessKline := &kls.Kline[i]
        accessKline.LowerBand = lowerBand[i]
    }
}

func(kl *Kline) IsClosedPriceIsBelowLowerBand() bool {
    if kl.ClosedPrice < kl.LowerBand {
        return true
    } else {
        return false
    } 
}

func(kl *Kline) IsClosedPriceIsAboveLowerBand() bool {
    if kl.ClosedPrice > kl.LowerBand {
        return true
    } else {
        return false
    } 
}

func (kl *Kline) IsClosedPriceIsAboveUpperBand() bool {
    if kl.ClosedPrice > kl.UpperBand {
        return true
    } else {
        return false
    }
}

func (kl *Kline) IsClosedPriceIsBellowUpperBand() bool {
    if kl.ClosedPrice < kl.UpperBand {
        return true
    } else {
        return false
    }
}
