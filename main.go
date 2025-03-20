package main

import (
	"fmt"
	"github.com/danielleedottech/danielleetech-go/handlers"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	RunServer()
}

func RunServer() {
	var err error

	if os.Getenv("ENV") == "production" {
		fmt.Println("running in port 80")
		err = http.ListenAndServe(":80", nil)
	} else {
		fmt.Println("running in port 8080")
		err = http.ListenAndServe(":8080", nil)
	}

	if err != nil {
		fmt.Println(err)
	}
}
