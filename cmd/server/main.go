package main

import (
	"log"

	"github.com/rjNemo/go-pass-gen/api"
)

const port = ":8080"

func main() {
	log.Printf("Start passgen server on http://localhost%s\n", port)
	log.Fatal(api.New().Start(port))
}
