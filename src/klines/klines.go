package klines

import (
    "context"
    "fmt"
    "time"
    "github.com/adshao/go-binance/v2"
    "github.com/b-quentin/indicator"
)

type Klines struct {
    Kline []Kline
    High []float64
    Low []float64
    Closed []float64
}

func timespanToDate(dateUnix int64) time.Time {
    return time.Unix(dateUnix/1000, 0)
}

func KlineBySymbolAndTime(kline *binance.WsKline, symbol string, interval string) (*binance.WsKline){
    if(kline.Symbol == symbol && kline.Interval == interval) {
        return kline
    } else {
        return nil
    }
}

func (kls *Klines) ShowLastKline() {
    kls.Kline[len(kls.Kline)-1].ShowKline()
}

func (kls *Klines) GetLastKline() Kline {
    return kls.Kline[len(kls.Kline)-1]
}

func (kls *Klines) GetPreviousKline() Kline {
    return kls.Kline[len(kls.Kline)-2]
}

func (kls *Klines) GetKlineForLeadingSpan() Kline {
    return kls.Kline[len(kls.Kline)-25]
}

func (kls *Klines) ShowKlines() {
    for _, kline := range kls.Kline {
        kline.ShowKline()
    }
}

func (kls *Klines) SetKlinesOutlimit(client *binance.Client, symbol string, interval string,startTime time.Time, endTime time.Time) {
    plage := 24 * time.Hour
    var allKlines []*binance.Kline
    // Récupérer les klines par plage de 8 heures
    for startTime.Before(endTime) {
        // Calculer la fin de la plage de 8 heuresmodify 
        plageEndTime := startTime.Add(plage)
        
        // Vérifier si la fin de la plage dépasse la date de fin
        if plageEndTime.After(endTime) {
            plageEndTime = endTime
        }
        
        klines, err := client.NewKlinesService().
            Symbol(symbol).
            Interval(interval).
            StartTime(startTime.Unix() * 1000).
            EndTime(plageEndTime.Unix() * 1000).
            Limit(1000). // Utilisez une limite supérieure pour obtenir plus de klines si nécessaire
            Do(context.Background())
        
        var newKlines []*binance.Kline
        if len(klines) > 0 {
            newKlines = klines[1:]
        }
        
        if err != nil {
        	fmt.Println("Erreur lors de la récupération des klines :", err)
        	return
        }
        
        // Ajouter les klines de la plage actuelle à allKlines
        allKlines = append(allKlines, newKlines...)
        
        // Afficher les klines de la plage actuelle
        //for _, kline := range klines {
        //	fmt.Printf("OpenTime: %s, CloseTime: %s, Open: %s, Close: %s\n",
        //		time.Unix(kline.OpenTime/1000, 0).UTC().Format(time.RFC3339),
        //		time.Unix(kline.CloseTime/1000, 0).UTC().Format(time.RFC3339),modify 
        //		kline.Open,
        //		kline.Close)
        //}
        
        // Mettre à jour le début de la plage pour la prochaine itération
        startTime = plageEndTime
}

    for _, k := range allKlines {
        openTime := timespanToDate(k.OpenTime)
        closedTime := timespanToDate(k.CloseTime)
        openValue := stringToFloat(k.Open)
        closeValue := stringToFloat(k.Close)
        highValue := stringToFloat(k.High)
        lowValue := stringToFloat(k.Low)
      
        kline := Kline{
            OpenTime: openTime,
            ClosedTime: closedTime,
            OpenPrice: openValue,
            ClosedPrice: closeValue,
            HighPrice: highValue,
            LowPrice: lowValue,
        }

        kls.Kline = append(kls.Kline, kline)
        kls.High = append(kls.High, highValue)
        kls.Low = append(kls.Low, lowValue)
        kls.Closed = append(kls.Closed, closeValue)
    }
}

func (kls *Klines) SetKlines(client *binance.Client, symbol string, interval string, limit int) {
    var klines []*binance.Kline 

    var err error
    if (limit == 0) {
        klines, err = client.NewKlinesService().Symbol(symbol).
            Interval(interval).Do(context.Background())
    } else {
        klines, err = client.NewKlinesService().Symbol(symbol).
            Interval(interval).Limit(limit).Do(context.Background())
    }

    if err != nil {
        fmt.Println(err)
    }

    for _, k := range klines {
        openTime := timespanToDate(k.OpenTime)
        closedTime := timespanToDate(k.CloseTime)
        openValue := stringToFloat(k.Open)
        closeValue := stringToFloat(k.Close)
        highValue := stringToFloat(k.High)
        lowValue := stringToFloat(k.Low)
      
        kline := Kline{
            OpenTime: openTime,
            ClosedTime: closedTime,
            OpenPrice: openValue,
            ClosedPrice: closeValue,
            HighPrice: highValue,
            LowPrice: lowValue,
        }

        kls.Kline = append(kls.Kline, kline)
        kls.High = append(kls.High, highValue)
        kls.Low = append(kls.Low, lowValue)
        kls.Closed = append(kls.Closed, closeValue)
    }
}

func (kls *Klines) SetKline(k *binance.WsKline) {
    kl := Kline{}
    kl.New(k)

    kls.Kline = append(kls.Kline, kl)
    kls.High = append(kls.High, kl.HighPrice)
    kls.Low = append(kls.Low, kl.LowPrice)
    kls.Closed = append(kls.Closed, kl.ClosedPrice)
}

func (kls *Klines) SetAtrSL(period int, multiplicateur float64) {
    _, atr := indicator.Atr(14, kls.High, kls.Low, kls.Closed)

    for i := range kls.Kline {
        var payback float64
        accessKline := &kls.Kline[i]
        atrWithMulti := atr[i] * multiplicateur

        if accessKline.ClosedPrice < accessKline.OpenPrice {
            payback = accessKline.ClosedPrice - atrWithMulti
        } else {
            payback = accessKline.OpenPrice - atrWithMulti
        }

        accessKline.AtrLower = payback
    }
}

func (kls *Klines) SetADX(period int) {
	adx, plus, minus := calculateADX(kls.Closed, kls.High, kls.Low, period)

    for i := range kls.Kline {
        accessKline := &kls.Kline[i]
        accessKline.Adx = adx[i]
        accessKline.Plus = plus[i]
        accessKline.Minus = minus[i]
    }
}

func (kls *Klines) ShowKlinesMRC() {
    for _, kline := range kls.Kline {
        kline.ShowMRC()
    }
}


