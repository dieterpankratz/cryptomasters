package main

import (
	"dpan/cryptomasters/api"
	"fmt"
)

func main() {
	rate, err := api.GetRate("ETH")
	if err == nil {
		fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
	}
}
