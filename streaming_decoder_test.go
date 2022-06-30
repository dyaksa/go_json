package go_json

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestDecoderReader(t *testing.T) {
	file, _ := os.Open("customer.json")
	decoder := json.NewDecoder(file)

	customer := &User{}

	err := decoder.Decode(customer)
	if err != nil {
		panic(err)
	}

	fmt.Print(customer)
}
