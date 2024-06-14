package tests

import (
	"ar/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestAnalyzeEndpoint(t *testing.T) {

	distinctOptions := make([]model.OptionsContract, 4)
	distinctOptions[0] = model.OptionsContract{
		Type:           "call",
		StrikePrice:    100.0,
		Bid:            1.0,
		Ask:            1.5,
		ExpirationDate: time.Now().AddDate(0, 0, 1),
		LongShort:      "long",
	}
	distinctOptions[1] = model.OptionsContract{
		Type:           "put",
		StrikePrice:    150.0,
		Bid:            1.5,
		Ask:            2.0,
		ExpirationDate: time.Now().AddDate(0, 0, 2),
		LongShort:      "short",
	}
	distinctOptions[2] = model.OptionsContract{
		Type:           "call",
		StrikePrice:    250.0,
		Bid:            2.5,
		Ask:            3.0,
		ExpirationDate: time.Now().AddDate(0, 0, 3),
		LongShort:      "long",
	}
	distinctOptions[3] = model.OptionsContract{
		Type:           "put",
		StrikePrice:    300.0,
		Bid:            3.0,
		Ask:            3.5,
		ExpirationDate: time.Now().AddDate(0, 0, 4),
		LongShort:      "short",
	}

	jsonData, err := json.Marshal(distinctOptions)
	if err != nil {
		t.Fatal(err)
	}

	reader := bytes.NewReader(jsonData)

	resp, err := http.Post("http://localhost:8080/analyze", "application/json", reader)
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	var result model.AnalysisResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)
}
