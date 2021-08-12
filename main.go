package main

import (
	"flag"

	"github.com/tgmendes/hello-argo/app"
)

func main() {
	servicePtr := flag.String("service", "svc-hello", "service to run")
	flag.Parse()

	if *servicePtr == "svc-hello" {
		app.RunHello(*servicePtr, ":8080")
	} else {
		app.RunCarValue(*servicePtr, ":8080")
	}
}
