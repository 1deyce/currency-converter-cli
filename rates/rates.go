package rates

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}

func FetchRates(apiURL string) (ExchangeRates, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return ExchangeRates{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ExchangeRates{}, errors.New("HTTP request failed with status code: " + resp.Status)
	}

	var rates ExchangeRates
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return ExchangeRates{}, err
	}
	// fmt.Printf("Fetched exchange rates: %+v\n", rates)

	return rates, nil
}