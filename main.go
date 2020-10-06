package main

import (
	"fmt"
	"net/http"
	"runtime"
)

const port = "8080"

var version = "v0.0.1"

func main() {
	http.HandleFunc("/hello", helloWorldHandler)
	fmt.Printf("Running on %v architecture, app version %v\n", runtime.GOARCH, version)
	fmt.Println("Listening on port", port)
	http.ListenAndServe(":"+port, nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received incoming web request %v on path %v\n", r.Method, r.URL.Path)
	w.Write([]byte("Hello version: " + version))
}
