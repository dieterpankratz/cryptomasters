package api

import (
	"dpan/cryptomasters/datatypes"
	"fmt"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/EUR"

func GetRate(currency string) (datatypes.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
}
