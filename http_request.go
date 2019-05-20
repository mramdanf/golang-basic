package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"net/url"
	"errors"
	"io/ioutil"
	"bytes"
)

type Student struct {
	ID string
	Name string
	Grade int
}

var BASE_URL = "http://localhost:8080/"

func fetchUsers() ([]Student, error) {
	var err error
	var client = &http.Client{}

	var students []Student

	request, err := http.NewRequest("POST", BASE_URL+"users", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status + string(bodyBytes))
	}

	err = json.Unmarshal(bodyBytes, &students)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func fetchUser(ID string) (Student, error) {
	var err error
	var client = &http.Client{}

	var params = url.Values{}
	params.Set("id", ID)
	var payload = bytes.NewBufferString(params.Encode())

	var student Student

	request, err := http.NewRequest("POST", BASE_URL+"user", payload)
	if err != nil {
		return student, err
	}
	request.Header.Set("content-type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return student, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return student, errors.New(response.Status + string(bodyBytes))
	}

	err = json.Unmarshal(bodyBytes, &student)
	if err != nil {
		return student, err
	}

	return student, nil
}

func main() {
	var students, err = fetchUsers()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	for _, student := range students {
		fmt.Printf("ID: %s \t Name: %s \t Grade: %d \n", student.ID, student.Name, student.Grade)
	}

	var ID = "W001"
	student, err := fetchUser(ID)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return
	}

	fmt.Printf("ID: %s \t Name: %s \t Grade: %d \n", student.ID, student.Name, student.Grade)
}