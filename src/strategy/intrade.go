package strategy

import "sync"

type I_InTrade interface {
	GetInTrade() bool
	SetInTrade(bool)
	GetSymbol() string
	SetSymbol(string)
    GetTakeProfit() float64
    SetTakeProfit(float64)
    GetStopLoss() float64
    SetStopLoss(float64)
    GetValueAtBuy() float64
    SetValueAtBuy(float64)
    GetPriceAtBuy() float64
    SetPriceAtBuy(float64)
}

type InTrade struct {
	InTrade bool
    Symbol string
    TakeProfit float64
    StopLoss float64
    ValueAtBuy float64
    PriceAtBuy float64
	mu   sync.Mutex
}

func (t *InTrade) GetInTrade() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.InTrade
}

func (t *InTrade) SetInTrade(value bool) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.InTrade = value
}

func (t *InTrade) GetSymbol() string {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.Symbol
}

func (t *InTrade) SetSymbol(value string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Symbol = value
}

func (t *InTrade) GetTakeProfit() float64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.TakeProfit
}

func (t *InTrade) SetTakeProfit(value float64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.TakeProfit = value
}

func (t *InTrade) GetStopLoss() float64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.StopLoss
}

func (t *InTrade) SetStopLoss(value float64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.StopLoss = value
}

func (t *InTrade) GetValueAtBuy() float64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.ValueAtBuy
}

func (t *InTrade) SetValueAtBuy(value float64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.ValueAtBuy = value
}


func (t *InTrade) GetPriceAtBuy() float64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.PriceAtBuy
}

func (t *InTrade) SetPriceAtBuy(value float64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.PriceAtBuy = value
}



