package main

import (
	"io"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "foo Ran")
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8000", nil)
}
