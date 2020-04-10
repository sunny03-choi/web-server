package main

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	mem := Member{"Sunwoo", 18, true}
	jsonBytes, err := json.Marshal(mem)
	if err != nil {
		panic(err)
	}
	jsonString := string(jsonBytes)
	fmt.Println(jsonString)
}
