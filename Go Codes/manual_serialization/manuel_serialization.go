package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Employee struct {
	Name string
	Role string
}

// Serialize: struct dizisini byte stream'e çevirdik

func serialize(employees []Employee) []byte {
	buf := new(bytes.Buffer)

	// kaç satır var (örneğin 3)
	binary.Write(buf, binary.BigEndian, int32(len(employees)))

	for _, e := range employees {
		writeString(buf, e.Name)
		writeString(buf, e.Role)
	}

	return buf.Bytes()
}

// Deserialize: byte stream'den struct dizisine çevirdik

func deserialize(data []byte) []Employee {
	buf := bytes.NewBuffer(data)

	var count int32
	binary.Read(buf, binary.BigEndian, &count)

	result := make([]Employee, count)
	for i := 0; i < int(count); i++ {
		name := readString(buf)
		role := readString(buf)
		result[i] = Employee{Name: name, Role: role}
	}

	return result
}

func writeString(buf *bytes.Buffer, s string) {
	length := int32(len(s))
	binary.Write(buf, binary.BigEndian, length)
	buf.WriteString(s)
}

func readString(buf *bytes.Buffer) string {
	var length int32
	binary.Read(buf, binary.BigEndian, &length)
	data := make([]byte, length)
	buf.Read(data)
	return string(data)
}

func main() {
	employees := []Employee{
		{"Fred", "Programmer"},
		{"Liping", "Analyst"},
		{"Sureerat", "Manager"},
	}

	serialized := serialize(employees)
	fmt.Println("Serialized Bytes:")
	fmt.Println(serialized)

	decoded := deserialize(serialized)
	fmt.Println("\nDeserialized Structs:")
	for _, e := range decoded {
		fmt.Printf("%-8s -> %s\n", e.Name, e.Role)
	}
}
