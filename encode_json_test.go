package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func LogJSON(data interface{}) {
	byte, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(byte))
}

func TestEncodeJSON(t *testing.T) {
	LogJSON("Dyaksa")
	LogJSON(12)
	LogJSON([]string{"Nama", "Domisili"})
}
