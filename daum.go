package main

import (
	"encoding/xml"
	"fmt"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	mem := Member{"Alex", 18, true}
	xmlBytes, err := xml.Marshal(mem)
	if err != nil {
		panic(err)
	}
	xmlstring := string(xmlBytes)
	fmt.Println(xmlstring)
}
