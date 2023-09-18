package klines

import (
	"math"

)

// Fonction pour calculer l'ADX, +DI et -DI
func calculateADX(closePrices, highPrices, lowPrices []float64, period int) ([]float64, []float64, []float64) {
	tr := make([]float64, len(closePrices))
	plusDM := make([]float64, len(closePrices))
	minusDM := make([]float64, len(closePrices))
	plusDI := make([]float64, len(closePrices))
	minusDI := make([]float64, len(closePrices))
	dx := make([]float64, len(closePrices))
	adx := make([]float64, len(closePrices))
	plusDIsma := 0.0
	minusDIsma := 0.0

	// Calcul du True Range (TR)
	for i := 1; i < len(closePrices); i++ {
		highLowDiff := math.Abs(highPrices[i] - lowPrices[i])
		highCloseDiff := math.Abs(highPrices[i] - closePrices[i-1])
		lowCloseDiff := math.Abs(lowPrices[i] - closePrices[i-1])
		tr[i] = math.Max(highLowDiff, math.Max(highCloseDiff, lowCloseDiff))
	}

	// Calcul du Plus Directional Movement (+DM) et du Minus Directional Movement (-DM)
	for i := 1; i < len(closePrices); i++ {
		plusDM[i] = math.Max(0, highPrices[i]-highPrices[i-1])
		minusDM[i] = math.Max(0, lowPrices[i-1]-lowPrices[i])
	}

	// Calcul du Plus Directional Indicator (+DI) et du Minus Directional Indicator (-DI)
	for i := period; i < len(closePrices); i++ {
		sumTR := 0.0
		sumPlusDM := 0.0
		sumMinusDM := 0.0

		for j := i - period + 1; j <= i; j++ {
			sumTR += tr[j]
			sumPlusDM += plusDM[j]
			sumMinusDM += minusDM[j]
		}

		plusDI[i] = (sumPlusDM / sumTR) * 100
		minusDI[i] = (sumMinusDM / sumTR) * 100
	}

	// Calcul de l'Average Directional Index (ADX)
	for i := period; i < len(closePrices); i++ {
		plusDIsma = (plusDIsma*(float64(period)-1) + plusDI[i]) / float64(period)
		minusDIsma = (minusDIsma*(float64(period)-1) + minusDI[i]) / float64(period)
		dx[i] = (math.Abs(plusDI[i]-minusDI[i]) / (plusDI[i] + minusDI[i])) * 100
		adx[i] = (adx[i-1]*(float64(period)-1) + dx[i]) / float64(period)
	}

	return adx, plusDI, minusDI
}
