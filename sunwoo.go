package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func main() {
	reqBody := bytes.NewBufferString("Post Plain text")
	resp, err := http.Post("http://superman125.dothome.co.kr/", "text/plain", reqBody)
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
