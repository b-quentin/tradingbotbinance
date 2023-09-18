package klines

import (

	"github.com/cinar/indicator"
)


func (kls Klines) DonchianTrend(period int) {
    highChannel, _, _ := indicator.DonchianChannel(period, kls.High)
    _, _, lowestChannel := indicator.DonchianChannel(period, kls.Low)


    //fmt.Println(highChannel)
    //fmt.Println(lowestChannel)
    //fmt.Println(kls.Closed)

    var trend []int
    var bull bool
    var bear bool

    first := true
    for i, cl := range kls.Closed {
        if i < period {
            trend = append(trend, 0)
            continue
        }

        if first {
            if kls.Closed[i-1] < highChannel[i-1] && cl > highChannel[i-1] {
                trend = append(trend, 1)
                bull = true
                first = false
            } else if kls.Closed[i-1] > lowestChannel[i-1] && cl < lowestChannel[i-1] {
                trend = append(trend, -1)
                bear = true
                first = false
            } else {
                trend = append(trend, 0)
            }
        }

        if bull {
            if kls.Closed[i-1] > lowestChannel[i-1] && cl < lowestChannel[i-1] {
                trend = append(trend, -1)
                bull = false
                bear = true
            } else {
                trend = append(trend, 1)
            }
        }

        if bear {
            if kls.Closed[i-1] < highChannel[i-1] && cl > highChannel[i-1] {
                trend = append(trend, 1)
                bear = false
                bull = true
            } else {
                trend = append(trend, -1)
            }
        }
    }

    //for i, test := range kls.Closed {
    //    fmt.Println("---------------------------------------")
    //    fmt.Printf("Closed:  %d | %f\n", i, test)
    //    fmt.Printf("High: %d | %f\n", i, highChannel[i])
    //    fmt.Printf("Low: %d | %f\n", i, lowestChannel[i])
    //    fmt.Printf("trend: %d\n", trend[i])

    //}

    for i := range kls.Kline {
        accessKline := &kls.Kline[i]
        accessKline.DonchianTrend = trend[i]   
    }
}


func (kl Kline) IsDonchianTrendBull() bool {
    if kl.DonchianTrend == 1 {
        return true
    } else {
        return false
    }
}
