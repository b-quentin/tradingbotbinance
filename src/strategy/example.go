package strategy

// Example
// BUY:
// ICHIMOKU
// Price > IsPriceAboveCloud
// Cross:
// Conversion line > Base line
// StockRSI Uptrend cross
// SELL:
// ATR 1.5 for stopLoss
// takeProfit 1.5

import (
	"fmt"
	"strconv"
	"tradingBot/src/klines"
	"tradingBot/src/logging"
	"tradingBot/src/order"
	"tradingBot/src/tools"

	//"test/src/order"

	"github.com/adshao/go-binance/v2"
	"github.com/google/uuid"
)

type Strategy4 struct {
    client *binance.Client
    id string
    symbol string
    price float64
    isTrade InTrade
    kls klines.Klines
    lastKl klines.Kline
    previousKl klines.Kline
    seekSell bool
    takeProfit float64
    stopLoss float64
    priceAtBuy float64
    valueAtBuy float64
    benef float64
    perte float64
    tx float64
}

func generateUniqueID() string {
    id := uuid.New()
    return id.String()
}

func (st *Strategy4) Launch(client *binance.Client,SYMBOL string, INTERVAL string, isTrade *InTrade, logging logging.I_Logging) {
    st.id = generateUniqueID()
    st.kls.SetKlines(client, SYMBOL, INTERVAL, 201)
    st.client = client
    st.kls.SetEma1(50)
    st.kls.SetEma2(200)
    st.kls.SetIchimoku()
    st.kls.SetAtrSL(14, 1.5)

    st.symbol = SYMBOL
    st.seekSell = false

    logging.LogData("[%s][%s][START]", st.id, st.symbol)

    st.symbolAlreadyIntrade()

    wsKlineHandler := func(event *binance.WsKlineEvent) {
        bool5m, kl := tools.KlineBySymbolAndTime(&event.Kline,SYMBOL,INTERVAL)
        bool1m, kl1m := tools.KlineBySymbolAndTime(&event.Kline,SYMBOL,"1m")

        if bool1m && kl1m.IsFinal { 
            st.price, _ = strconv.ParseFloat(kl1m.Close, 64) 
            logging.LogDataPrice("[%s][%s] price: %f", st.id, st.symbol, st.price)
        }

        if bool5m && kl.IsFinal {
            st.refreshData(kl)

            logging.LogData("[%s][%s] previous: %v", st.id, st.symbol, st.previousKl.GetKline())
            logging.LogData("[%s][%s] lastKL: %v", st.id, st.symbol, st.lastKl.GetKline())
            logging.LogData("[%s][%s] checkEma: %t", st.id, st.symbol, st.lastKl.IsEma1AboveEma2())
            logging.LogData("[%s][%s] checkIchimokuIsClosedPriceAboveLeadingSpanA: %t", st.id, st.symbol, st.lastKl.IsClosedPriceAboveLeadingSpanA())
            logging.LogData("[%s][%s] checkIchimokuIsClosedPriceAboveLeadingSpanB: %t", st.id, st.symbol, st.lastKl.IsClosedPriceAboveLeadingSpanB())
            logging.LogData("[%s][%s] checkIchimokuConversionAboveBase: %t", st.id, st.symbol, st.lastKl.IsConversionAboveBase())
            logging.LogData("[%s][%s] checkStockRsiPreviousKlIsDownTrend: %t", st.id, st.symbol, st.lastKl.StockRsi.IsDownTrend())
            logging.LogData("[%s][%s] checkStockRsiLastKlUptrend: %t", st.id, st.symbol, st.lastKl.StockRsi.IsUpTrend())
            logging.LogData("[%s][%s] crossSignal: %t", st.id, st.symbol, st.crossSignal())

            st.buy()

            logging.LogData("[%s][%s] SeekSell: %t", st.id, st.symbol, st.seekSell)
        }

        if bool1m && kl1m.IsFinal { 
            st.sellWin()
            st.sellLoose()
        }
    }

    errHandler := func(err error) {
        fmt.Println(err)
    }

    symbol := map[string]string{SYMBOL: INTERVAL}
    symbol1m := map[string]string{SYMBOL: "1m"}
    doneSymbol, _, _ := binance.WsCombinedKlineServe(symbol, wsKlineHandler, errHandler)
    doneSymbol1m, _, _ := binance.WsCombinedKlineServe(symbol1m, wsKlineHandler, errHandler)
    <-doneSymbol1m
    <-doneSymbol
}

func (st *Strategy4) symbolAlreadyIntrade() {
    if st.isTrade.GetInTrade() && st.isTrade.GetSymbol() == st.symbol {
        st.takeProfit = st.isTrade.GetTakeProfit()
        st.stopLoss = st.isTrade.GetStopLoss()
        st.valueAtBuy = st.isTrade.GetValueAtBuy()
        st.priceAtBuy = st.isTrade.GetPriceAtBuy()
        st.seekSell = true
    }
}

func (st *Strategy4) refreshData(kl *binance.WsKline) {
    st.kls.SetKline(kl)
    st.kls.SetEma1(50)
    st.kls.SetEma2(200)
    st.kls.SetKlinesStockRSI()
    st.kls.SetIchimoku()
    st.kls.SetAtrSL(14, 1.5)
    st.lastKl = st.kls.GetLastKline()
    st.previousKl = st.kls.GetPreviousKline()
}


//func (st *Strategy4) crossSignal() bool {
//    if st.previousKl.StockRsi.IsBelowLower() && 
//    st.previousKl.StockRsi.IsDownTrend() &&
//    st.lastKl.StockRsi.IsUpTrend() {
//        return true
//    } else {
//        return false
//    }
//
//}

func (st *Strategy4) crossSignal() bool {
    if st.previousKl.StockRsi.IsDownTrend() && 
    st.lastKl.StockRsi.IsUpTrend() {
        return true
    } else {
        return false
    }

}

func (st *Strategy4) checkEntry() bool {
    if !st.seekSell && 
        !st.isTrade.GetInTrade() && 
        st.lastKl.IsEma1AboveEma2() &&
        st.lastKl.IsClosedPriceAboveCloud() &&
        st.lastKl.IsConversionAboveBase() &&
        st.crossSignal() {
        return true
    } else {
        return false
    }
}

func (st *Strategy4) getUnitPrice(fills []*binance.Fill) {
    total := 0.00
    i := 0.00
    for _, fill := range fills {
        prix, _ := strconv.ParseFloat(fill.Price, 64)
        total = total + prix
        i = i + 1.00
    }

    st.priceAtBuy = total / i
}

func (st *Strategy4) buy() {
    if st.checkEntry() {
        st.isTrade.SetInTrade(true)

        st.seekSell = true
        // ACHAT
        resp, err := order.BuyMarketOrder(st.client, st.symbol)

        if err != nil {
            logging.LogError("[%s][ERROR] %s", st.id, err)
        } else {
            st.valueAtBuy, _ = strconv.ParseFloat(resp.CummulativeQuoteQuantity, 64)
            st.getUnitPrice(resp.Fills)
            //Take Profit = Prix d'entrée + (Ratio x (Prix d'entrée - Stop Loss))
            st.stopLoss = st.lastKl.AtrLower 
            st.takeProfit = st.lastKl.ClosedPrice + (1.5 * (st.priceAtBuy - st.stopLoss))

            st.isTrade.SetTakeProfit(st.takeProfit)
            st.isTrade.SetStopLoss(st.stopLoss)

            st.isTrade.SetSymbol(st.symbol)
            st.isTrade.SetPriceAtBuy(st.priceAtBuy)
            st.isTrade.SetValueAtBuy(st.valueAtBuy)

            logging.LogData("[%s][%s][BUY] price: %f takeProfit: %f stopLoss: %f trade: %v", st.id, st.symbol, st.priceAtBuy, st.takeProfit, st.stopLoss, resp)
            logging.LogTrade("[%s][%s][BUY] price: %f takeProfit: %f stopLoss: %f trade: %v", st.id, st.symbol, st.priceAtBuy, st.takeProfit, st.stopLoss, resp)
        }
    }
}

func (st *Strategy4) checkTakeProfit() bool {
    // note st.takeProfit ce retouve parfois à zero je ne sais pas pourquoi donc pour éviter de prendre une vente je le block par un control
    // Peux être du à plusieurs achats en même temps ?
    if st.isTrade.GetSymbol() == st.symbol && st.seekSell && st.takeProfit < st.lastKl.ClosedPrice && st.takeProfit > 0.000000 {
        return true
    } else {
        return false
    }
}

func (st *Strategy4) checkStopLoss() bool {
    if st.isTrade.GetSymbol() == st.symbol && st.seekSell && st.stopLoss > st.lastKl.ClosedPrice && st.stopLoss > 0.000000 {
        return true
    } else {
        return false
    }
}


func (st *Strategy4) reset() {
    st.seekSell = false
    st.isTrade.SetInTrade(false)
}

func (st *Strategy4) calculBenef(valueAtSell float64) {
    if st.isTrade.GetInTrade() && st.isTrade.GetSymbol() == st.symbol {
        st.valueAtBuy = st.isTrade.GetValueAtBuy() 
    }

    st.benef = valueAtSell - st.valueAtBuy
}

func (st *Strategy4) calculTauxEvol(priceSell float64) {
    st.tx = ((priceSell - st.priceAtBuy) / priceSell)
}

func (st *Strategy4) getUnitPriceSell(fills []*binance.Fill) float64 {
    total := 0.00
    i := 0.00
    for _, fill := range fills {
        prix, _ := strconv.ParseFloat(fill.Price, 64)
        total = total + prix
        i = i + 1.00
    }

    return total / i
}

func (st *Strategy4) sellWin() {
    if st.checkTakeProfit() {
        logging.LogTrade("[%s][%s][ENTRY-SELL][WIN]", st.id, st.symbol)
        // SELL WIN
        resp, err := order.SellMarketOrder(st.client, st.symbol)

        if err != nil {
            logging.LogError("[%s][ERROR] %s", st.id, err)
        } else {
            valueAtSell, _ := strconv.ParseFloat(resp.CummulativeQuoteQuantity, 64)

            unitePriceSell := st.getUnitPriceSell(resp.Fills)
            st.calculBenef(valueAtSell)
            st.calculTauxEvol(unitePriceSell)

            logging.LogData("[%s][%s][SELL][WIN] price: %f profit: %f taux: %f trade: %v", st.id, st.symbol, unitePriceSell, st.benef, st.tx, resp)
            logging.LogTrade("[%s][%s][SELL][WIN] price: %f profit: %f taux: %f trade: %v", st.id, st.symbol, unitePriceSell, st.benef, st.tx, resp)
            st.reset()
        }
    }   
}

func (st *Strategy4) calculPerte(valueAtSell float64) {
    if st.isTrade.GetInTrade() && st.isTrade.GetSymbol() == st.symbol {
        st.valueAtBuy = st.isTrade.GetValueAtBuy() 
    }

    st.perte = st.valueAtBuy - valueAtSell
}

func (st *Strategy4) sellLoose() {
    if st.checkStopLoss() {
        logging.LogTrade("[%s][%s][ENTRY-SELL][WIN]", st.id, st.symbol)
        // SELL WIN
        resp, err := order.SellMarketOrder(st.client, st.symbol)

        if err != nil {
            logging.LogError("[%s][ERROR] %s", st.id, err)
        } else {
            valueAtSell, _ := strconv.ParseFloat(resp.CummulativeQuoteQuantity, 64)

            unitePriceSell := st.getUnitPriceSell(resp.Fills)
            st.calculPerte(valueAtSell)
            st.calculTauxEvol(unitePriceSell)

            logging.LogData("[%s][%s][SELL][LOOSE] price: %f, perte: %f, taux: %f, trade: %v", st.id, st.symbol, unitePriceSell, st.perte, st.tx, resp)
            logging.LogTrade("[%s][%s][SELL][LOOSE] price: %f, perte: %f, taux: %f, trade: %v", st.id, st.symbol, unitePriceSell, st.perte, st.tx, resp)
            st.reset()
    }
    }   
}

func (st *Strategy4) getUnitPriceAndAmount(sellOrder *binance.CreateOCOResponse) (float64, float64) {
    total := 0.00
    i := 0.00
    for _, fill := range sellOrder.OrderReports {
        prix, _ := strconv.ParseFloat(fill.CummulativeQuoteQuantity, 64)
        total = total + prix
        i = i + 1.00
    }

    return total , total / i
}
