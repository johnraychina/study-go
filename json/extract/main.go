package extract

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"remark":"cold_data"}`

	var data map[string]string
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for key, value := range data {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

}
