package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	PostalCode string
}

type User struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Address    []Address
}

func TestJSONObject(t *testing.T) {
	customer := User{
		FirstName:  "Dyaksa",
		MiddleName: "Jauharuddin",
		LastName:   "Nour",
		Age:        27,
		Married:    false,
	}

	byte, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}
