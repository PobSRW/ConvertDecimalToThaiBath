package main

import (
	"fmt"
	"sorawat-convert-currency-suffix/service"

	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}

	result, err := service.ConvertCurrency(inputs)

	fmt.Println(result, "result")
	fmt.Println(err, "err")
}
