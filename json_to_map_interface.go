package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var jsonString = `{"Name": "Ramdan", "Age": 24}`
	var jsonData = []byte(jsonString)

	var user map[string]interface{}
	var err = json.Unmarshal(jsonData, &user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Fullname: ", user["Name"])
	fmt.Println("Age: ", user["Age"])
}