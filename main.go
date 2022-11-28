package main

import (
	"encoding/json"
	"fmt"
)

type SensorReading struct {
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Time        string `json:"time"`
	Information info   `json:"info"`
}
type info struct {
	Description string `json:"desc"`
}
type Book struct {
	Author Author `json:"author"`
	Title  string `json:"title"`
}
type Author struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Developer bool   `json:"is_developer"`
}

func main() {
	//	fmt.Println("Hello, World!")
	author := Author{Name: "me", Age: 34, Developer: true}
	book := Book{Author: author, Title: "Learning concurrency"}
	fmt.Printf("%v\n", book)
	byteArray, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteArray))

	jsonString := `{ "name":"batterysaver" , "capacity":40 , "time":"2022" , "info": {
		"desc" : "densor reading"
	}}`
	var read SensorReading
	error := json.Unmarshal([]byte(jsonString), &read)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(read)
}
