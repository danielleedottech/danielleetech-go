package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	RunServer()
}

func RunServer() {
	var err error

	if os.Getenv("ENV") == "production" {
		err = http.ListenAndServe(":80", nil)
	} else {
		err = http.ListenAndServe(":8080", nil)
	}

	if err != nil {
		fmt.Println(err)
	}
}
