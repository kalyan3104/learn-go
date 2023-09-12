// Basic web server
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
		name := req.URL.Query().Get("name")
		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
	})
	http.ListenAndServe(":8080", nil)
}

// http://localhost:8080/hello?name=YourName
