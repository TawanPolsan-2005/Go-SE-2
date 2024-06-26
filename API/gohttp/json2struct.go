package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int 
	Name string 
	Age  int 
}

func main() {
	data := []byte(`{
		"id": 2,
		"name": "polsan",
		"age": 19
	}`)

	var u User
	err := json.Unmarshal(data, &u)

	fmt.Printf("% #v\n", u)
	fmt.Println(err)
}