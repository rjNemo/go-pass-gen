package main

import (
	"net/http"

	"github.com/rjNemo/go-pass-gen/api"
	"github.com/rjNemo/go-pass-gen/cmd"
)

func main() {
	//serveWeb()
	cli()
}

func cli() {
	cmd.Execute()
}

func serveWeb() {
	s := api.NewServer()
	http.ListenAndServe(":8080", s)
}
