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

func RunCarValue(cmdFlag string, port string) {
	log.Println("Botting up car value app...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{model}", func(w http.ResponseWriter, r *http.Request) {
		carValue(w, r, cmdFlag)
	})

	log.Printf("Listening on port %s", port)
	_ = http.ListenAndServe(port, r)
}

func carValue(w http.ResponseWriter, r *http.Request, cmdFlag string) {
	commitHash := os.Getenv("COMMIT_HASH")
	model := chi.URLParam(r, "model")

	rand.Seed(time.Now().Unix())
	lowerValue := 500
	upperValue := 100000
	value := rand.Intn(upperValue-lowerValue) + lowerValue

	carMsg := `<h1>🚗 Car Valuation Service 🏎</h1>
<hr \>
<p>🤑 Based on our calculations, this is what your %s is worth: £%d 🎉</p>

<p>📱 Running on the following commit hash: %s.</p>
`
	carMsgB := []byte(fmt.Sprintf(carMsg, model, value, commitHash))
	_, _ = w.Write(carMsgB)
}
