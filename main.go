package main

import (
	"flag"
	"github.com/tgmendes/hello-argo/app"
)

func main() {
	servicePtr := flag.String("service", "hello", "service to run")
	flag.Parse()

	if *servicePtr == "hello" {
		app.RunHello()
	} else {
		app.RunCarValue()
	}
}
