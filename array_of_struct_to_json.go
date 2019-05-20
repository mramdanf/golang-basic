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
	var users = [...]User{
		{"Ramdan", 24},
		{"Lisda Adistiani", 22},
	}
	var jsonData, err = json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}