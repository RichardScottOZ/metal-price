package data

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestNewRates(t *testing.T) {
	rs, err := NewRates(log.New(os.Stdout, "", log.LstdFlags))

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", rs.rate)
}
