package capital

import (
	"fmt"
)

type Capital struct {
    Capital float64
    Crypto float64
    Win int64
    Loose int64
}

func (Cp *Capital) Init(capital float64) {
    Cp.Capital = capital
    Cp.Win = 0
    Cp.Loose = 0
}

func (Cp Capital) ShowCapital() {
    fmt.Println("Capital: ", Cp.Capital)
}

func (Cp Capital) ShowCrypto() {
    fmt.Println("Crypto:", Cp.Crypto)
}

func (cp *Capital) SetCapital(capital float64) {
    cp.Capital = capital
}

func (cp *Capital) SetCrypto(price float64) {
    crypto := cp.Capital / price
    cp.Crypto = crypto - (crypto * 0.001)

    cp.Capital = cp.Capital - (cp.Crypto * price) 
}

func (cp *Capital) SetWin(price float64) {
    cp.Capital = cp.Capital + (cp.Crypto * price)
    cp.Win = cp.Win + 1
}

func (cp *Capital) SetLoose(price float64) {
    cp.Capital = cp.Capital + (cp.Crypto * price)
    cp.Loose = cp.Loose + 1
}

func (cp Capital) ShowWin() {
    fmt.Println("Win:", cp.Win)
}

func (cp Capital) ShowLoose() {
    fmt.Println("Loose:", cp.Loose)
}

