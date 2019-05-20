package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	var jsonString = `[
		{ "Name": "Ramdan", "Age": 24 },
		{ "Name": "Lisda Adistiani", "Age": 22 }
	]`

	var jsonData = []byte(jsonString)

	var users []User
	var err = json.Unmarshal(jsonData, &users)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, user := range users {
		fmt.Println("FullName: ", user.FullName)
		fmt.Println("Age: ", user.Age)
		fmt.Println()
	}
}