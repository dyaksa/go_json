package go_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonMapDecode(t *testing.T) {
	jsonStr := `{"id": 1, "name": "Cod", "type": true}`
	jsonByte := []byte(jsonStr)

	var result map[string]interface{}

	err := json.Unmarshal(jsonByte, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])

}

func TestJsonMapEncode(t *testing.T) {
	result := map[string]interface{}{
		"id":   1,
		"name": "Dyaksa Jauharuddin",
		"type": false,
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
