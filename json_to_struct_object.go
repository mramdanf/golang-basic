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
	var jsonString = `{"Name": "Ramdan", "Age": 24}`
	var jsonData = []byte(jsonString)

	var user User
	var err = json.Unmarshal(jsonData, &user)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Fullname: ", user.FullName)
	fmt.Println("Age: ", user.Age)
}