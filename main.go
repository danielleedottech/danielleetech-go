package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/danielleedottech/danielleetech-go/templates"
)

func main() {
	http.Handle("/", templ.Handler(templates.Hello("World")))
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
