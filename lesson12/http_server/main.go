package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed,error:", err)
	}
}

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintf(w, "%v\n", r.Form)
	fmt.Fprintf(w, "path:%s\n", r.URL.Path)
	fmt.Fprintf(w, "schema:%s\n", r.URL.Scheme)
	fmt.Fprintf(w, "hello world")
}
