package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	RunServer()
}

func RunServer() {
	var err error

	if os.Getenv("ENV") == "production" {
		fmt.Println("running in port 80")
		err = http.ListenAndServe(":80", nil)
	} else {
		fmt.Println("running in port 80")
		err = http.ListenAndServe(":8080", nil)
	}

	if err != nil {
		fmt.Println(err)
	}
}
