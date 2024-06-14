package controllers

import (
	"ar/model"
	"testing"
	"time"
)

func TestAnalyzeOptionsContractsValidation(t *testing.T) {

	// Valid contracts
	validContracts := []model.OptionsContract{
		{
			Type:           model.CallOptionType,
			StrikePrice:    100.0,
			Bid:            1.0,
			Ask:            1.5,
			ExpirationDate: time.Now().AddDate(0, 0, 1),
			LongShort:      model.Long,
		},
		{
			Type:           model.PutOptionType,
			StrikePrice:    150.0,
			Bid:            1.5,
			Ask:            2.0,
			ExpirationDate: time.Now().AddDate(0, 0, 2),
			LongShort:      model.Short,
		},
	}

	_, err := AnalyzeOptionsContracts(validContracts)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	// Invalid contracts
	invalidContracts := []model.OptionsContract{
		{
			Type:           "invalid",
			StrikePrice:    100.0,
			Bid:            1.0,
			Ask:            1.5,
			ExpirationDate: time.Now().AddDate(0, 0, 1),
			LongShort:      model.Long,
		},
		{
			Type:           model.CallOptionType,
			StrikePrice:    150.0,
			Bid:            2.0,
			Ask:            1.5,
			ExpirationDate: time.Now().AddDate(0, 0, -1),
			LongShort:      model.Short,
		},
		{
			Type:           model.PutOptionType,
			StrikePrice:    -100.0,
			Bid:            1.5,
			Ask:            2.0,
			ExpirationDate: time.Now().AddDate(0, 0, 2),
			LongShort:      model.Long,
		},
	}

	for _, contract := range invalidContracts {
		_, err := AnalyzeOptionsContracts([]model.OptionsContract{contract})
		if err == nil {
			t.Errorf("Expected error, got nil for contract: %+v", contract)
		}
	}

}

func TestAnalyzeOptionsContracts(t *testing.T) {

	contracts := []model.OptionsContract{
		{Type: model.CallOptionType, StrikePrice: 50, Bid: 2, Ask: 3, LongShort: model.Long, ExpirationDate: time.Now().AddDate(0, 0, 1)},
		{Type: model.PutOptionType, StrikePrice: 60, Bid: 3, Ask: 4, LongShort: model.Short, ExpirationDate: time.Now().AddDate(0, 0, 1)},
		{Type: model.CallOptionType, StrikePrice: 70, Bid: 4, Ask: 5, LongShort: model.Long, ExpirationDate: time.Now().AddDate(0, 0, 1)},
	}

	result, err := AnalyzeOptionsContracts(contracts)
	if err != nil {
		t.Errorf("AnalyzeOptionsContracts returned an error: %v", err)
	}

	// Assert the expected values
	expectedXValues := []model.XYValue{
		{X: 0.0, Y: -63},
		{X: 1, Y: -62},
		{X: 2, Y: -61},
		{X: 3, Y: -60},
		{X: 4, Y: -59},
		{X: 5, Y: -58},
	}

	if !compareXYValues(result.XYValues[0:6], expectedXValues) {
		t.Errorf("Unexpected XYValues. Got %v, want %v", result.XYValues, expectedXValues)
	}

	expectedMaxProfit := 1877.0
	if result.MaxProfit != expectedMaxProfit {
		t.Errorf("Unexpected MaxProfit. Got %f, want %f", result.MaxProfit, expectedMaxProfit)
	}

	expectedMaxLoss := -63.0
	if result.MaxLoss != expectedMaxLoss {
		t.Errorf("Unexpected MaxLoss. Got %f, want %f", result.MaxLoss, expectedMaxLoss)
	}

	expectedBreakEvenPoints := []float64{}
	if !compareFloat64Slices(result.BreakEvenPoints, expectedBreakEvenPoints) {
		t.Errorf("Unexpected BreakEvenPoints. Got %v, want %v", result.BreakEvenPoints, expectedBreakEvenPoints)
	}
}

func compareXYValues(a, b []model.XYValue) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func compareFloat64Slices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
