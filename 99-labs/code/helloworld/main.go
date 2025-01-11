package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

var hostname string
var version string

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from %s running Go version %s!\n", hostname, version)
}

func main() {
	h, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	v := runtime.Version()
	version = v
	hostname = h
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
