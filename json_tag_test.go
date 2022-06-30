package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:     "1",
		Name:   "Gembong",
		Amount: "10000",
	}

	jsonByte, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonByte))
}

func TestDecodeJsonTag(t *testing.T) {
	jsonStr := `{"ID":"1","NAME":"Gembong","AMOUNT":"10000"}`
	jsonByte := []byte(jsonStr)

	product := &Product{}

	err := json.Unmarshal(jsonByte, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}
