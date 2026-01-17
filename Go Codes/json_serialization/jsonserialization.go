package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family string
	Person string
}

type Email struct {
	Kind    string
	Address string
}

// JSON Kaydetme

func saveJSON(filename string, key interface{}) {
	data, err := json.MarshalIndent(key, "", "   ")
	checkError(err)

	err = os.WriteFile(filename, data, 0600)
	checkError(err)

	fmt.Println("JSON saved to", filename)
}

func loadJSON(filename string, key interface{}) {
	data, err := os.ReadFile(filename)
	checkError(err)

	err = json.Unmarshal(data, key)
	checkError(err)

	fmt.Println(" JSON loaded successfully from", filename)
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
