package second

import (
	"fmt"
	"testing"
)

func TestChromedp(t *testing.T) {
	id := 1
	name := "hahah"
	requestBody := fmt.Sprintf(`{
	"id":%d",
	"name": "%s"
	}`, id, name)

	type RequestBody struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}

	var jsonStr = []byte(requestBody)
	t.Logf("%s", jsonStr)
}
