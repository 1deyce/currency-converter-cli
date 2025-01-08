package main

import (
	"encoding/json"
	"fmt"

	// "log"
	"net/http"
	"os"
	"strconv"

	"github.com/1deyce/currency-converter/converter"
	"github.com/1deyce/currency-converter/rates"

	// "github.com/charmbracelet/huh"
	// "github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

type CurrencyData struct {
	From 	string  `json:"from"`
	To   	string  `json:"to"`
	Amount 	string `json:"amount"`
} 

type ConversionResponse struct {
	ConvertedAmount float64 `json:"amount"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var data CurrencyData

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

    err = json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	from := data.From
	to := data.To
	amountStr := data.Amount

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

	response := ConversionResponse {
		ConvertedAmount: converted,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)
    if err != nil {
        fmt.Printf("Error encoding response: %v\n", err)
    }
}

func main() {
	// app := fiber.New()

    // app.Get("/", func(c fiber.Ctx) error {
    //     return c.SendString("Hello, World 👋!")
    // })

    // log.Fatal(app.Listen(":8000"))
	http.HandleFunc("/convert", handler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}	


