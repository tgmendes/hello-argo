package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func RunCarValue() {
	commitHash := os.Getenv("COMMIT_HASH")
	log.Println("Botting up car value app...")
	log.Printf("Commit hash: %s\n", commitHash)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{model}", carValue)

	log.Println("Listening on port :8080")
	_ = http.ListenAndServe(":8080", r)
}

func carValue(w http.ResponseWriter, r *http.Request) {
	model := chi.URLParam(r, "model")

	rand.Seed(time.Now().Unix())
	lowerValue := 500
	upperValue := 100000
	value := rand.Intn(upperValue-lowerValue) + lowerValue
	resp := fmt.Sprintf("Your car %s is worth £%d", model, value)
	_, _ = w.Write([]byte(resp))
}
