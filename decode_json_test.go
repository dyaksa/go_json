package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDecodeJSON(t *testing.T) {
	jsonstr := `{"FirstName": "Dyaksa", "MiddleName": "Jauharuddin", "LastName": "Nour", "Age": 27, "Married": false}`
	jsonEncode := []byte(jsonstr)

	customer := &User{}

	err := json.Unmarshal(jsonEncode, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
}
