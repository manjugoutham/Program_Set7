package main

import (
	"encoding/json"
	"fmt"
)

type Friend struct {
	Name           string `json:"name"`
	MobileNumber   string `json:"mobileNumber"`
	Address        string `json:"address"`
	CityOrTownName string `json:"cityOrTownName"`
}

func main() {
	friends := []Friend{
		{Name: "Goutham", MobileNumber: "1234567890", Address: "#456 banglore Main road", CityOrTownName: "banglore"},
		{Name: "Smith", MobileNumber: "9876543211", Address: "#456 banglore kr puram", CityOrTownName: "banglore"},
		{Name: "Rahul", MobileNumber: "9876543221", Address: "456 banglore kr puram", CityOrTownName: "kolar"},
		{Name: "Ashwath", MobileNumber: "9876598765", Address: "#456 banglore Main road", CityOrTownName: "kolar"},
	}

	// Convert the slice to JSON
	friendsJSON, err := json.Marshal(friends)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(friendsJSON))
}
