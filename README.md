# Currency Converter

A simple currency converter application that fetches exchange rates and allows users to convert amounts between different currencies.

## Features

- Fetches real-time exchange rates from Open Exchange Rates API.
- Converts amounts between selected currencies.
- Validates user input for currency selection and amount.

## Packages Used

- net/http for http requests to the currency exchange api
- github.com/charmbracelet/huh for the TUI interface form
- encoding/json in order to marshal the data for the api

## Requirements

- Go (version 1.14 or later)
- An API key from [Open Exchange Rates]

## Example

Please select the currency you wish to convert FROM:
1. United States Dollar
2. Great Britain Pound
3. Euro
4. Japanese Yen
5. South African Rand

Please enter the amount you wish to convert:
? 100

Please select the currency you wish to convert TO:
1. United States Dollar
2. Great Britain Pound
3. Euro
4. Japanese Yen
5. South African Rand

Converted amount: 85.00 GBP