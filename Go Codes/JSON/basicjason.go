package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Email string
}

func main() {
	p := Person{Name: "Alice", Email: "alice@example.com"}

	// Go -> JSON
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	// JSON -> Go
	var decode Person
	json.Unmarshal(data, &decode)
	fmt.Println(decode)
}
