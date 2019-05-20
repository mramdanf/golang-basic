package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	FullName string
	Age int
}

func main() {
	var user = User{"Ramdan", 24}

	var jsonData, err = json.Marshal(user)
	if err != nil {
		fmt.Println(err.Error())
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}