package klines

import (

	"github.com/cinar/indicator"
)

func (kls *Klines) SetIchimoku() {
    conversionBase , baseLigne, leadingSpanA, leadingSpanB, laggingSpan := indicator.IchimokuCloud(kls.High, kls.Low, kls.Closed)

    for i := range kls.Kline {
        accessKline := &kls.Kline[i]
        accessKline.ConversionBase = conversionBase[i]
        accessKline.BaseLigne = baseLigne[i]
        if i > 24 {
            accessKline.LeadingSpanA = leadingSpanA[i-25]
            accessKline.LeadingSpanB = leadingSpanB[i-25]
        }
        accessKline.LaggingSpan = laggingSpan[i]
    }
}

func (kl *Kline) IsClosedPriceAboveCloud() bool {
    if kl.ClosedPrice > kl.LeadingSpanA && kl.ClosedPrice > kl.LeadingSpanB {
        return true
    } else {
        return false
    }
}

func (kl *Kline) IsClosedPriceBellowCloud() bool {
    if kl.ClosedPrice < kl.LeadingSpanA && kl.ClosedPrice < kl.LeadingSpanB {
        return true
    } else {
        return false
    }
}

func (kl *Kline) IsClosedPriceAboveLeadingSpanA() bool {
    if kl.ClosedPrice > kl.LeadingSpanA {
        return true
    } else {
        return false
    }
}

func (kl *Kline) IsClosedPriceAboveLeadingSpanB() bool {
    if kl.ClosedPrice > kl.LeadingSpanB {
        return true
    } else {
        return false
    }
}

func (kl *Kline) IsClosedPriceBellowLeadingSpanA() bool {
    if kl.ClosedPrice < kl.LeadingSpanA {
        return true
    } else {
        return false
    }
}

func (kl *Kline) IsClosedPriceBellowLeadingSpanB() bool {
    if kl.ClosedPrice < kl.LeadingSpanB {
        return true
    } else {
        return false
    }
}

func (kl *Kline) IsConversionAboveBase() bool {
    if kl.ConversionBase > kl.BaseLigne {
        return true
    } else {
        return false
    }
}

func (kl *Kline) IsConversionBellowBase() bool {
    if kl.ConversionBase > kl.BaseLigne {
        return true
    } else {
        return false
    }
}


