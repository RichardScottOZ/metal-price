package services

import (
	"fmt"
	"testing"

	metal "github.com/chutified/metal-price/metal/service/protos/metal"
	"google.golang.org/grpc"
	"gopkg.in/go-playground/assert.v1"
)

func TestMetal(t *testing.T) {

	metalConn, err := grpc.Dial("localhsot:10552", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to dial localhost:10521: %v", err)
	}
	client := metal.NewMetalClient(metalConn)
	ms := NewMetal(client)

	assert.Equal(t, ms.client, client)

	tests := []struct {
		name      string
		metal     string
		expErrMsg string
	}{
		// {
		//     name:      "ok",
		//     base:      "USD",
		//     dest:      "EUR",
		//     expErrMsg: "",
		// },
		{
			name:      "service not running - periodic symbol",
			metal:     "au",
			expErrMsg: "metal service",
		},
		{
			name:      "invalid base",
			metal:     "invalid",
			expErrMsg: "material .* not found",
		},
		{
			name:      "service not running - periodic symbol",
			metal:     "au",
			expErrMsg: "metal service",
		},
		{
			name:      "service not running",
			metal:     "gold",
			expErrMsg: "metal service",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			p, err := ms.GetPrice(test.metal)
			if err != nil {

				exp := fmt.Sprintf(".*%s.*", test.expErrMsg)
				assert.MatchRegex(t1, err.Error(), exp)

			} else {

				assert.Equal(t1, "", test.expErrMsg)
				assert.NotEqual(t1, p, 0)
			}
		})
	}
}
