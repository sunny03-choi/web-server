package main

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

//Person -
type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"sunwoo", 18}
	pbytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	resp, err := http.Post("http://supermnan.dothome.co.kr", "application/xml", buff)
}
