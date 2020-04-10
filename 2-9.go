package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{"sunwoo", 18}
	pbytes, _ := xml.Marshal(person)
	buff := bytes.NewBuffer(pbytes)
	req, err := http.NewRequest("POST", "http;//httpbin.org/post", buff)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/xml")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		str := string(respBody)
		println(str)
	}
}
