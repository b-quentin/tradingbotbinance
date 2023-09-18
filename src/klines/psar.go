package klines

import (
	"github.com/cinar/indicator"
)

func (kls *Klines) ParabolicSar() {
    if len(kls.Closed) == 0 {
        return
    }
    psar, dir := indicator.ParabolicSar(kls.High, kls.Low, kls.Closed)

    for i := range kls.Kline {
        accessKline := &kls.Kline[i]
        accessKline.Psar = psar[i]
        accessKline.PsarTrend = int(dir[i])   
    }
}


func (kl *Kline) IsPSARBullTrend() bool {
    if kl.PsarTrend == 1 {
        return true
    } else {
        return false
    }
}

func (kl *Kline) IsPSARBearTrend() bool {
    if kl.PsarTrend == 1 {
        return true
    } else {
        return false
    }
}
