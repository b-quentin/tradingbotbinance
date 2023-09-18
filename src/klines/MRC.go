package klines

import (
	"math"
)

type MRC struct {
    // SuperSmoother function
	MeanLine float64

	// Calculate upper and lower bands
	UpBand1 float64
	LoBand1 float64
	UpBand2 float64
	LoBand2 float64

    UpBand2_1 float64
    LoBand2_1 float64
    UpBand2_2 float64 
    LoBand2_2 float64 
    UpBand2_3 float64 
    LoBand2_3 float64 
    UpBand2_4 float64 
    LoBand2_4 float64 
    UpBand2_5 float64 
    LoBand2_5 float64 
    UpBand2_6 float64 
    LoBand2_6 float64 
    UpBand2_7 float64 
    LoBand2_7 float64 
    UpBand2_8 float64 
    LoBand2_8 float64 
    UpBand2_9 float64 
    LoBand2_9 float64 
} 

func (kls *Klines) supersmoother(length int) []float64 {
    src := kls.Closed
	s_a1 := math.Exp(-math.Sqrt(2) * math.Pi / float64(length))
	s_b1 := 2 * s_a1 * math.Cos(math.Sqrt(2)*math.Pi/float64(length))
	s_c3 := -math.Pow(s_a1, 2)
	s_c2 := s_b1
	s_c1 := 1 - s_c2 - s_c3
	ss := make([]float64, len(src))
	for i := range src {
		if i == 0 {
			ss[i] = s_c1 * src[i]
		} else if i == 1 {
			ss[i] = s_c1*src[i] + s_c2*ss[i-1]
		} else {
			ss[i] = s_c1*src[i] + s_c2*ss[i-1] + s_c3*ss[i-2]
		}
	}
	return ss
}

func (kls *Klines) supersmootherTR(src []float64,length int) []float64 {
	s_a1 := math.Exp(-math.Sqrt(2) * math.Pi / float64(length))
	s_b1 := 2 * s_a1 * math.Cos(math.Sqrt(2)*math.Pi/float64(length))
	s_c3 := -math.Pow(s_a1, 2)
	s_c2 := s_b1
	s_c1 := 1 - s_c2 - s_c3
	ss := make([]float64, len(src))
	for i := range src {
		if i == 0 {
			ss[i] = s_c1 * src[i]
		} else if i == 1 {
			ss[i] = s_c1*src[i] + s_c2*ss[i-1]
		} else {
			ss[i] = s_c1*src[i] + s_c2*ss[i-1] + s_c3*ss[i-2]
		}
	}
	return ss
}

func (kls *Klines) calculTR() []float64 {
	tr := make([]float64, len(kls.Closed))
    for i := range kls.Kline {
        kl := &kls.Kline[i]
        tr[i] = math.Max(kl.HighPrice-kl.LowPrice, math.Max(kl.HighPrice-kl.ClosedPrice, kl.ClosedPrice-kl.LowPrice))
    }

    return tr
}

// 200 , 1, 2.145 
func (kls *Klines) MeanReversionChannel(length int, innerMult, outerMult float64) {
	pi := 2 * math.Asin(1)
	mult := pi * innerMult
	mult2 := pi * outerMult
	gradSize := 0.5

	// SuperSmoother function
	meanLine := kls.supersmoother(length)

    tr := kls.calculTR()
    meanRangeArray := kls.supersmootherTR(tr, length)

    for i := range kls.Kline {
        kl := &kls.Kline[i]
		kl.MRC.MeanLine = meanLine[i]
		meanRange := meanRangeArray[i]

		kl.MRC.UpBand1 = meanLine[i] + (meanRange*mult)
		kl.MRC.LoBand1 = meanLine[i] - (meanRange*mult)
		kl.MRC.UpBand2 = meanLine[i] + (meanRange*mult2)
		kl.MRC.LoBand2 = meanLine[i] - (meanRange*mult2)

        kl.MRC.UpBand2_1 = kl.MRC.UpBand2 + (meanRange * gradSize * 4) 
        kl.MRC.LoBand2_1 = kl.MRC.LoBand2 - (meanRange * gradSize * 4) 

        kl.MRC.UpBand2_2 = kl.MRC.UpBand2 + (meanRange * gradSize * 3) 
        kl.MRC.LoBand2_2 = kl.MRC.LoBand2 - (meanRange * gradSize * 3) 

        kl.MRC.UpBand2_3 = kl.MRC.UpBand2 + (meanRange * gradSize * 2) 
        kl.MRC.LoBand2_3 = kl.MRC.LoBand2 - (meanRange * gradSize * 2) 

        kl.MRC.UpBand2_4 = kl.MRC.UpBand2 + (meanRange * gradSize * 1) 
        kl.MRC.LoBand2_4 = kl.MRC.LoBand2 - (meanRange * gradSize * 1) 

        kl.MRC.UpBand2_5 = kl.MRC.UpBand2 + (meanRange * gradSize * 0) 
        kl.MRC.LoBand2_5 = kl.MRC.LoBand2 - (meanRange * gradSize * 0) 

        kl.MRC.UpBand2_6 = kl.MRC.UpBand2 + (meanRange * gradSize * -1)
        kl.MRC.LoBand2_6 = kl.MRC.LoBand2 - (meanRange * gradSize * -1)

        kl.MRC.UpBand2_7 = kl.MRC.UpBand2 + (meanRange * gradSize * -2)
        kl.MRC.LoBand2_7 = kl.MRC.LoBand2 - (meanRange * gradSize * -2)

        kl.MRC.UpBand2_8 = kl.MRC.UpBand2 + (meanRange * gradSize * -3)
        kl.MRC.LoBand2_8 = kl.MRC.LoBand2 - (meanRange * gradSize * -3)

        kl.MRC.UpBand2_9 = kl.MRC.UpBand2 + (meanRange * gradSize * -4)
        kl.MRC.LoBand2_9 = kl.MRC.LoBand2 - (meanRange * gradSize * -4)
	}
}

func(kl *Kline) PriceAboveMeanLine () bool {
    if kl.MRC.MeanLine < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowMeanLine () bool {
    if kl.MRC.MeanLine > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveUpBand1 () bool {
    if kl.MRC.UpBand1 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand1 () bool {
    if kl.MRC.UpBand1 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveUpBand2 () bool {
    if kl.MRC.UpBand2 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand2 () bool {
    if kl.MRC.UpBand2 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveLoBand1 () bool {
    if kl.MRC.LoBand1 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowLoBand1 () bool {
    if kl.MRC.LoBand1 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveLoBand2 () bool {
    if kl.MRC.LoBand2 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowLoBand2 () bool {
    if kl.MRC.LoBand2 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}
func(kl *Kline) PriceAboveUpBand2_1 () bool {
    if kl.MRC.UpBand2_1 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand2_1 () bool {
    if kl.MRC.UpBand2_1 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveUpBand2_2 () bool {
    if kl.MRC.UpBand2_2 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand2_2 () bool {
    if kl.MRC.UpBand2_2 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveUpBand2_3 () bool {
    if kl.MRC.UpBand2_3 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand2_3 () bool {
    if kl.MRC.UpBand2_3 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveUpBand2_4 () bool {
    if kl.MRC.UpBand2_4 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand2_4 () bool {
    if kl.MRC.UpBand2_4 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveUpBand2_5 () bool {
    if kl.MRC.UpBand2_5 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand2_5 () bool {
    if kl.MRC.UpBand2_5 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveUpBand2_6 () bool {
    if kl.MRC.UpBand2_6 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand2_6 () bool {
    if kl.MRC.UpBand2_6 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveUpBand2_7 () bool {
    if kl.MRC.UpBand2_7 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand2_7 () bool {
    if kl.MRC.UpBand2_7 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveUpBand2_8 () bool {
    if kl.MRC.UpBand2_8 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand2_8 () bool {
    if kl.MRC.UpBand2_8 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceAboveUpBand2_9 () bool {
    if kl.MRC.UpBand2_9 < kl.ClosedPrice {
        return true
    } else {
        return false
    }
}

func(kl *Kline) PriceBellowUpBand2_9 () bool {
    if kl.MRC.UpBand2_9 > kl.ClosedPrice {
        return true
    } else {
        return false
    }
}


