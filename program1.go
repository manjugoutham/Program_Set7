package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	stateCapital := map[string]string{
		"Gujarat":        "Gandhinagar",
		"Tamil Nadu":     "Chennai",
		"Andhra Pradesh": "Amaravati",
		"Maharashtra":    "Mumbai",
		"Karnataka":      "Bengaluru",
	}

	fmt.Println("stateCapital:", stateCapital)

	stateJSON, err := json.Marshal(stateCapital)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(stateJSON))
}
