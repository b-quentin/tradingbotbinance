package klines

import "github.com/cinar/indicator"

func (kls *Klines) SetMACD() {
    macd, signal := indicator.Macd(kls.Closed)

    for i := range kls.Kline {
        accessKline := &kls.Kline[i]
        accessKline.Macd = macd[i]    
        accessKline.Signal = signal[i]    
    }

}

func(kl *Kline) IsMacdAndSignalBelowZero() bool {
    if kl.Macd < 0 && kl.Signal < 0 {
        return true
    } else {
        return false
    } 
}

func(kl *Kline) IsMacdBelowSignal() bool {
    if kl.Macd < kl.Signal {
        return true
    } else {
        return false
    }
}

func(kl *Kline) IsMacdAboveSignal() bool {
    if kl.Macd > kl.Signal {
        return true
    } else {
        return false
    }
}
