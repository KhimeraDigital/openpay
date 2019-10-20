# OpenPay API Client
[Based in fairbank-io/openpay](https://github.com/fairbank-io/openpay)

[![Software License](https://img.shields.io/badge/license-MIT-red.svg)](LICENSE)

Pure Go [OpenPay](https://www.openpay.mx/) client implementation.

## Example

```go
package main

import (
	"gitlab.com/khimera.digital/public/openpay-go/openpay"
	// If you're not using go modules, use this import instead
	// openpay "gitlab.com/khimera.digital/public/openpay-go/openpay-go"
)

func main() {
	// Start a new client instance
	client, err := openpay.NewClient("API_KEY", "MERCHANT_ID", nil)
	if err != nil {
		panic(err)
	}
	// Register customer
	rick := openpay.Customer{
		Name:     "Rick",
		LastName: "Sanchez",
		Email:    "rick@mail.com",
		Address: &openpay.Address{
			CountryCode: "MX",
			PostalCode:  "94560",
		},
	}
	client.Customers.Create(&rick)

	// Add Card
	card := &openpay.Card{
		HolderName:      "Rick Sanchez",
		CardNumber:      "4111111111111111",
		CVV2:            "401",
		ExpirationMonth: "10",
		ExpirationYear:  "19",
		Address:         rick.Address,
	}
	client.Charges.AddCard(card)

	// Execute charge
	sale := &openpay.ChargeWithStoredCard{
		Charge: openpay.Charge{
			Method:      "card",
			Amount:      1000,
			Currency:    "MXN",
			Description: "sample charge operation",
			Customer:    rick,
		},
		SourceID: card.ID,
		Capture:  true,
	}
	client.Charges.WithCard(sale)
}
```
