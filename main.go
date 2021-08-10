package main

import (
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	servicePtr := flag.String("service", "hello", "service to run")
	flag.Parse()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	if *servicePtr == "hello" {
		r.Get("/", runHello)
	} else {
		r.Get("/", runRaffle)
	}
	log.Println("Listening on port :8080")
	_ = http.ListenAndServe(":8080", r)
}

func runHello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("name")
	if name == "" {
		name = "Cuvva"
	}
	_, _ = w.Write([]byte(fmt.Sprintf("Hello %s!", name)))
	return
}

func runRaffle(w http.ResponseWriter, r *http.Request) {
	n := rand.Int()
	_, _ = w.Write([]byte(fmt.Sprintf("your lucky number is: %d", n)))
	return
}
