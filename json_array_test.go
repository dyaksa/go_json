package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodeJSONArray(t *testing.T) {
	customer := User{
		FirstName:  "Dyaksa",
		MiddleName: "Jauharuddin",
		LastName:   "Nour",
		Age:        27,
		Married:    false,
		Hobbies:    []string{"Code", "Hiking", "Gym"},
	}

	data, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func TestDecodeJSONArray(t *testing.T) {
	jsonStr := `{"FirstName":"Dyaksa","MiddleName":"Jauharuddin","LastName":"Nour","Age":27,"Married":false,"Hobbies":["Code","Hiking","Gym"]}`
	jsonByte := []byte(jsonStr)

	customer := &User{}

	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := User{
		FirstName: "Dyakksa",
		Address: []Address{
			{
				Street:     "Jalan Tidak Ditemnukan",
				PostalCode: "59515",
			},
			{
				Street:     "Kenep",
				PostalCode: "59572",
			},
		},
	}
	jsonByte, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonByte))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonStr := `{"FirstName":"Dyakksa","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Address":[{"Street":"Jalan Tidak Ditemnukan","PostalCode":"59515"},{"Street":"Kenep","PostalCode":"59572"}]}`
	jsonByte := []byte(jsonStr)

	customer := &User{}

	err := json.Unmarshal(jsonByte, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.Address)
	fmt.Println(customer.FirstName)
}

func TestOnlyJSONArrayComplex(t *testing.T) {
	jsonStr := `[{"Street":"Jalan Tidak Ditemnukan","PostalCode":"59515"},{"Street":"Kenep","PostalCode":"59572"}]`
	jsonByte := []byte(jsonStr)

	address := &[]Address{}

	err := json.Unmarshal(jsonByte, address)
	if err != nil {
		panic(err)
	}

	fmt.Println(address)

	//for _, add := range *address {
	//	fmt.Println(add.PostalCode)
	//}

}

func TestJSONArrayComplexAddress(t *testing.T) {
	address := []Address{
		{
			Street:     "Jalan Mangunjiwan",
			PostalCode: "57677",
		},
		{
			Street:     "Jalan Pamulang",
			PostalCode: "7657676",
		},
	}

	jsonByte, err := json.Marshal(address)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonByte))

}
