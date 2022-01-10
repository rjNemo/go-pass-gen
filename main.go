package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/rjNemo/go-pass-gen/api"
	"github.com/rjNemo/go-pass-gen/cmd"
)

func main() {
	web := flag.Bool("web", false, "")
	flag.Parse()

	if *web {
		serveWeb()
		return
	}
	cli()
}

func cli() {
	cmd.Execute()
}

func serveWeb() {
	s := api.NewServer()
	log.Fatal(http.ListenAndServe(":8080", s))
}
