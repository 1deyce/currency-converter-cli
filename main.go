package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/1deyce/currency-converter/converter"
	"github.com/1deyce/currency-converter/rates"
	"github.com/charmbracelet/huh"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	var from, to, amountStr string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Please select the currency you wish to convert FROM").
				Options(
					huh.NewOption("United States Dollar", "USD"),
					huh.NewOption("Great Britain Pound", "GBP"),
					huh.NewOption("Euro", "EUR"),
					huh.NewOption("Japanese Yen", "JPY"),
					huh.NewOption("South African Rand", "ZAR"),
				).
				Value(&from),
		
			huh.NewInput().
				Title("Please enter the amount you wish to convert").
				Prompt("? ").
				Value(&amountStr),
		
			huh.NewSelect[string]().
				Title("Please select the currency you wish to convert TO").
				Options(
					huh.NewOption("United States Dollar", "USD"),
					huh.NewOption("Great Britain Pound", "GBP"),
					huh.NewOption("Euro", "EUR"),
					huh.NewOption("Japanese Yen", "JPY"),
					huh.NewOption("South African Rand", "ZAR"),
				).
				Value(&to),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
        fmt.Printf("Error parsing amount: %v\n", err)
        return
    }

	converted, err := converter.Convert(from, to, amount, rates)
	if err != nil {
        fmt.Printf("Error converting currency: %v\n", err)
        return
    }

	fmt.Printf("Converted amount: %.2f %s\n", converted, to)
}	


