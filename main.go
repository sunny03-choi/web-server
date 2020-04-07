package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", new(testHandler))
	log.Printf("### %s", "abcd")
	http.ListenAndServe(":5000", nil)
}

type testHandler struct {
	http.Handler
}

func (h *testHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	str := "Your Requset Path is" + req.URL.Path
	w.Write([]byte(str))
}

type Handler interface {
	ServeHttp(w http.ResponseWriter, r *http.Request)
}
