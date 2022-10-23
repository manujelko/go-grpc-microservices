package data

import (
	"testing"

	"github.com/hashicorp/go-hclog"
)

func TestNewExchangeRates(t *testing.T) {
	_, err := NewExchangeRates(hclog.Default())

	if err != nil {
		t.Fatal(err)
	}
}
