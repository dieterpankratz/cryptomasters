package api_test

import (
	"dpan/cryptomasters/api"
	"testing"
)

func TestXxx(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error was not found")
	}
}
