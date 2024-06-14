package controllers

import (
	"ar/model"
	"math"
	"time"
)

func AnalyzeOptionsContracts(contracts []model.OptionsContract) (model.AnalysisResult, error) {

	// Validate options contracts
	if err := validateOptionsContracts(contracts); err != nil {
		return model.AnalysisResult{}, err
	}

	var xValues []model.XYValue
	var maxProfit float64 = math.Inf(-1)
	var maxLoss float64 = math.Inf(1)
	var breakEvenPoints []float64

	// Loop through a range of possible underlying prices
	for price := 0.0; price <= 1000.0; price += 1 {
		totalProfitLoss := 0.0
		for _, contract := range contracts {
			// Calculate profit or loss for each contract at this price
			var profitOrLoss float64
			if contract.Type == model.CallOptionType {
				if price > contract.StrikePrice {
					// Profit = (Price - Strike Price) - Bid
					profitOrLoss = (price - contract.StrikePrice) - contract.Bid
				} else {
					// Loss = -Bid
					profitOrLoss = -contract.Bid
				}
			} else if contract.Type == model.PutOptionType {
				if price < contract.StrikePrice {
					// Profit = (Strike Price - Price) - Bid
					profitOrLoss = (contract.StrikePrice - price) - contract.Bid
				} else {
					// Loss = -Bid
					profitOrLoss = -contract.Bid
				}
			}

			// Adjust profit or loss based on long or short position
			if contract.LongShort == model.Short {
				profitOrLoss = -profitOrLoss
			}

			totalProfitLoss += profitOrLoss
		}

		// Append X & Y values
		xValues = append(xValues, model.XYValue{X: price, Y: totalProfitLoss})

		// Update maximum profit and maximum loss
		if totalProfitLoss > maxProfit {
			maxProfit = totalProfitLoss
		}
		if totalProfitLoss < maxLoss {
			maxLoss = totalProfitLoss
		}

		// Check for break-even points
		if equal(totalProfitLoss, 0) {
			breakEvenPoints = append(breakEvenPoints, price)
		}
	}

	return model.AnalysisResult{
		XYValues:        xValues,
		MaxProfit:       maxProfit,
		MaxLoss:         maxLoss,
		BreakEvenPoints: breakEvenPoints,
	}, nil
}

const epsilon = 1e-9

func equal(a, b float64) bool {
	return math.Abs(a-b) <= epsilon
}

func validateOptionsContracts(contracts []model.OptionsContract) error {
	for _, contract := range contracts {
		if contract.Type != model.CallOptionType && contract.Type != model.PutOptionType {
			return model.ErrOptionContractValidation
		}
		if contract.LongShort != model.Long && contract.LongShort != model.Short {
			return model.ErrOptionContractValidation
		}

		// Validate expiration date
		if contract.ExpirationDate.Before(time.Now()) {
			return model.ErrOptionContractValidation
		}

		// Validate bid and ask prices
		if contract.Bid < 0 || contract.Ask < 0 || contract.Bid > contract.Ask {
			return model.ErrOptionContractValidation
		}

		// Validate strike price
		if contract.StrikePrice <= 0 {
			return model.ErrOptionContractValidation
		}

	}
	return nil
}
