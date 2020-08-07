package main

import (
	"fmt"
	"net/http"
)

// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	// create response binary data
	greeting := greeting("Code.education Rocks!")
	data := []byte(greeting) // slice of bytes

	// write `data` to response
	res.Write(data)
}

func greeting(term string) string {
	return fmt.Sprintf("<b>%s</b>", term)
}

func main() {
	// create a new handler
	handler := HttpHandler{}

	// listen and serve
	http.ListenAndServe(":8000", handler)
}
