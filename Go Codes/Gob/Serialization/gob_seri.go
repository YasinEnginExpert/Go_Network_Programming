package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	person := Person{
		Name: Name{Family: "Yasins", Personal: "Yasin"},
		Email: []Email{
			{Kind: "home", Address: "Yasin@Yasins.name"},
			{Kind: "work", Address: "Yasinh@Yasins.edu.tr"},
		},
	}

	saveGob("Personal.gob", person)
}

// Serileştirme Fonksiyonu
func saveGob(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}



func checkError(err error) {
	if err != nil {
		log.Fatalln("Error:",err)
	}
}
