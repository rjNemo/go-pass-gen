package main

import (
	"flag"
	"log"

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
	log.Fatal(api.NewServer().Start(":8080"))
}
