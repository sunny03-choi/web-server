package main

import "io/ioutil"

func main() {
	bytes, err := ioutil.ReadFile("C:\\temp\\1.txt")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("C:\\temp\\2.txt", bytes, 0)
	if err != nil {
		panic(err)
	}
}
