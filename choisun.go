package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	resp, err := http.PostForm("http://superman.dothome.co.kr", url.Values{"Name": {"Choi"}, "Age": {"18"}})
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
