package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := `{
        "Thinly Sliced Peeled Apples" : "6 Cups",
        "sugar" : "3/4 cup",
        "flour" : "2 tablespoons",
        "cinamon" : "1/4 teaspoon",
        "nutmeg" : "1/8 tablespoon",
        "lemon juice": "1 tablespoon"
    }`

	// Convert the JSON to map
	var mapdata map[string]string
	err := json.Unmarshal([]byte(data), &mapdata)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for key, value := range mapdata {
		fmt.Printf("%s: %s\n", key, value)
	}
}
