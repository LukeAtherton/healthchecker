package main

import (
	"log"
	"net/http"

	. "github.com/lukeatherton/healthchecker/app"
)

func main() {
	if err := http.ListenAndServe(":8001", NewRouter()); err != nil {
		log.Fatal(err)
	}
}
