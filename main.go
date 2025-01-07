/*todo: Convert between a number of base currencies. Recommended to start with USD,GBP,EUR & JPY.
Use the huh package in order to create the terminal form
You'll need to make use of a third party API in order to obtain currency conversion data.*/

package main

import (
	"fmt"
	"os"
	"github.com/1deyce/currency-converter/rates"
	// "github.com/1deyce/currency-converter/converter"
)

func main() {
	// TODO: get 3 inputs from user [amount, from, to]
	// use options select for from & to string values using form
	/* var from, to string
    var amount float64*/

	appID := os.Getenv("OPENEXCHANGE_APP_ID")
	if appID == "" {
        fmt.Println("OPENEXCHANGE_APP_ID not set, please set it in your environment variables")
		return
    }

    apiURL := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s&symbols=USD,GBP,EUR,JPY,ZAR", appID)

	rates, err := rates.FetchRates(apiURL)
	if err != nil {
        fmt.Printf("Error fetching exchange rates: %v\n", err)
        return
    }

	fmt.Printf("Fetched rates: %+v\n", rates)

	// converted := converter.Convert(from, to, amount, rates)
}	


