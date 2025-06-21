package main

import (
	"dpan/cryptomasters/api"
	"fmt"
)

func main() {
	rate, err := api.GetRate("BTC")

	fmt.Println(rate.Price, err)
}
