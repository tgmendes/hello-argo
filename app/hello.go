package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func RunHello() {
	commitHash := os.Getenv("COMMIT_HASH")
	log.Println("Botting up hello app...")
	log.Printf("Commit hash: %s\n", commitHash)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", hello)

	log.Println("Listening on port :8080")
	_ = http.ListenAndServe(":8080", r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("name")
	if name == "" {
		name = "Cuvva"
	}
	_, _ = w.Write([]byte(fmt.Sprintf("Hello %s!", name)))
}
