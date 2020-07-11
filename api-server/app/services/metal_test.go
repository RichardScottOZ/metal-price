package services

import (
	"fmt"
	"testing"

	metal "github.com/chutified/metal-price/metal/service/protos/metal"
	"google.golang.org/grpc"
	"gopkg.in/go-playground/assert.v1"
)

func TestMetal(t *testing.T) {

	// >>>>>>>>>>>>>>> NewMetal
	metalConn, err := grpc.Dial("localhost:10502", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("unable to dial localhost:10501: %v", err)
	}
	defer metalConn.Close()
	client := metal.NewMetalClient(metalConn)
	ms := NewMetal(client)

	assert.Equal(t, ms.client, client)

	tests := []struct {
		name      string
		action    func()
		metal     string
		expErrMsg string
	}{
		{
			name:      "ok",
			action:    func() {},
			metal:     "silver",
			expErrMsg: "",
		},
		{
			name:      "ok periodic symbol",
			action:    func() {},
			metal:     "silver",
			expErrMsg: "",
		},
		{
			name:      "invalid metal",
			action:    func() {},
			metal:     "invalid",
			expErrMsg: "material .* not found",
		},
		{
			name:      "service not running periodic symbol",
			action:    func() { metalConn.Close() },
			metal:     "au",
			expErrMsg: "metal service",
		},
		{
			name:      "service not running",
			action:    func() { metalConn.Close() },
			metal:     "gold",
			expErrMsg: "metal service",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t1 *testing.T) {

			test.action()

			// >>>>>>>>>>>>>>> GetPrice
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
