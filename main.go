package main

import (
	"log"
	"net/http"

	"github.com/rjNemo/go-pass-gen/api"
	"github.com/rjNemo/go-pass-gen/cmd"
)

func main() {
	serveWeb()
	//cli()
}

func cli() {
	cmd.Execute()
}

func serveWeb() {
	s := api.NewServer()
	log.Fatal(http.ListenAndServe(":8080", s))
}
