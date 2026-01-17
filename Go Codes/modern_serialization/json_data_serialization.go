package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {

	Name string `json:"name"`
	Role string `json:"role"` // JSON alan adıyla eşleştirir

}

func main() {
	// Orijinal veri

	employees := []Employee{
		{Name: "Yasin", Role: "Programmer"},
		{Name: "Hakan", Role: "Analyst"},
		{Name: "Necmi", Role: "Manager"},
	}

	// Serialization(Marshalling)

	data, err := json.Marshal(employees) // Go struct'larını JSON formatına çevirir
	if err != nil {
		panic(err)
	}

	fmt.Println("Serialized (JSON)")
	fmt.Println(string(data))

	// Deseriazlization

	var decoded []Employee
	err = json.Unmarshal(data, &decoded) // JSON verisini tekrar Go struct'a dönüştürür
	if err != nil {
		panic(err)
	}

	fmt.Println("\n Deserialized (Struct)")
	for _, emp := range decoded {
		fmt.Printf("Name: %-8s | Role: %s\n",emp.Name, emp.Role)
	}

}