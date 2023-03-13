package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Friend1 struct {
	Name           string `json:"name"`
	MobileNumber   string `json:"mobileNumber"`
	Address        string `json:"address"`
	CityOrTownName string `json:"cityOrTownName"`
}

func main() {
	friendsdata, err := ioutil.ReadFile("friend.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var friends []Friend1
	err = json.Unmarshal(friendsdata, &friends)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(friends)
}
