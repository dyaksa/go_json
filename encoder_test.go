package go_json

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	file, _ := os.Create("customerout.json")
	encoder := json.NewEncoder(file)

	customer := User{
		FirstName:  "Dyaksa",
		MiddleName: "Jauharudin",
		LastName:   "Nour",
	}

	err := encoder.Encode(customer)
	if err != nil {
		panic(err)
	}
}
