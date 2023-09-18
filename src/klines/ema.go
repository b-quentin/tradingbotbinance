package klines

import "github.com/cinar/indicator"

func (kls *Klines) SetEma1(period int) {
    ema := indicator.Ema(period, kls.Closed)

    for i := range kls.Kline {
        accessKline := &kls.Kline[i]
        accessKline.Ema1 = ema[i]   
    }
}

func (kls *Klines) SetEma2(period int) {
    ema := indicator.Ema(period, kls.Closed)

    for i := range kls.Kline {
        accessKline := &kls.Kline[i]
        accessKline.Ema2 = ema[i]    
    }
}

func(kl *Kline) IsEma1AboveEma2() bool {
    if kl.Ema1 > kl.Ema2 {
        return true
    } else {
        return false
    } 
}

func(kl *Kline) IsEma2AboveEma1() bool {
    if kl.Ema2 > kl.Ema1 {
        return true
    } else {
        return false
    } 
}


