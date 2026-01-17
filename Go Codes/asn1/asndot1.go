package main

import (
	"encoding/asn1"
	"fmt"
	"log"
	"time"
)

type Person struct {
	Name string
	Age  int
}

// Integer
func exampleInt() {
	val := 13

	mdata, _ := asn1.Marshal(val) // ASN.1byte dizisine çevirir

	var n int
	asn1.Unmarshal(mdata, &n) // Go değerine çevirir
	fmt.Printf("[INTEGER]\nBefore: %v, After: %v\n\n", val, n)
}

// String
func exampleString() {
	s := "hello"
	mdata, _ := asn1.Marshal(s)
	var newStr string
	asn1.Unmarshal(mdata, &newStr)
	fmt.Printf("[STRING]\nBefore: %v, After: %v\n\n", s, newStr)
}

// Time
func exampleTime() {
	t := time.Now()
	mdata, _ := asn1.Marshal(t)
	var newTime time.Time
	asn1.Unmarshal(mdata, &newTime)
	fmt.Printf("[TIME]\nBefore: %v\nAfter:  %v\n\n", t, newTime)
}

func exampleStruct() {
	p1 := Person{"Alice", 30}
	mdata, _ := asn1.Marshal(p1)

	var p2 Person
	asn1.Unmarshal(mdata, &p2)
	fmt.Printf("[STRUCT]\nBefore: %+v\nAfter:  %+v\n\n", p1, p2)
}

// Error
type BadStruct struct {
	field1 int // küçük harfle başlıyor → exportable değil!
	field2 int
}

func exampleError() {
	b := BadStruct{1, 2}
	_, err := asn1.Marshal(b)
	if err != nil {
		log.Printf("[ERROR] Marshal failed: %v\n\n", err)
	}
}

func main() {

	exampleInt()
	exampleString()
	exampleTime()
	exampleStruct()
	exampleError()
}
