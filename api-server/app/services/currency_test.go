package services

import (
	"fmt"
	"testing"

	currency "github.com/chutommy/metal-price/currency/service/protos/currency"
	"google.golang.org/grpc"
	"gopkg.in/go-playground/assert.v1"
)

func TestCurrency(t *testing.T) {

	// >>>>>>>>>>>>>>> NewCurrency
	currencyConn, err := grpc.Dial("localhost:10501", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to dial localhost:10501: %v", err)
	}
	defer currencyConn.Close()
	client := currency.NewCurrencyClient(currencyConn)
	cs := NewCurrency(client)

	assert.Equal(t, cs.client, client)

	tests := []struct {
		name      string
		action    func()
		base      string
		dest      string
		expErrMsg string
	}{
		{
			name:      "ok",
			action:    func() {},
			base:      "USD",
			dest:      "EUR",
			expErrMsg: "",
		},
		{
			name:      "invalid base",
			action:    func() {},
			base:      "invalid",
			dest:      "EUR",
			expErrMsg: "base currency .* not found",
		},
		{
			name:      "invalid destination",
			action:    func() {},
			base:      "USD",
			dest:      "invalid",
			expErrMsg: "base currency .* not found",
		},
		{
			name:      "service not running",
			action:    func() { currencyConn.Close() },
			base:      "USD",
			dest:      "EUR",
			expErrMsg: "currency service",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			test.action()

			// >>>>>>>>>>>>>>> GetRate
			r, err := cs.GetRate(test.base, test.dest)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.expErrMsg)
				assert.NotEqual(t1, r, 0)
			}
		})
	}
}
