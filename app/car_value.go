package app

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RunCarValue(cmdFlag string, port string) {
	log.Println("Botting up car value app...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		carValue(w, r, cmdFlag)
	})

	r.Get("/{model}", func(w http.ResponseWriter, r *http.Request) {
		carValue(w, r, cmdFlag)
	})

	log.Printf("Listening on port %s", port)
	_ = http.ListenAndServe(port, r)
}

func carValue(w http.ResponseWriter, r *http.Request, cmdFlag string) {
	commitHash := os.Getenv("COMMIT_HASH")
	model := chi.URLParam(r, "model")

	if model == "" {
		model = "random"
	}

	rand.Seed(time.Now().Unix())
	lowerValue := 500
	upperValue := 100000
	value := rand.Intn(upperValue-lowerValue) + lowerValue
	gif := getCarGIF(model)

	carMsg := `<h1>🚗 Car Valuation service! 🏎</h1>
<div>
<iframe src="%s" frameborder="0" width="800"></iframe>
</div>
<hr \>
<p>🤑 Based on our calculations, this is what your <strong>%s</strong> is worth: <strong>£%d</strong> 🎉</p>

<p>📱 Running on the following commit hash: <strong>%s</strong>.</p>


`
	carMsgB := []byte(fmt.Sprintf(carMsg, gif, model, value, commitHash))
	_, _ = w.Write(carMsgB)
}

func getCarGIF(searchTerm string) string {
	cl := GiphyClient{
		APIKey: "UwelgUbg9VUspxHAf3E8XXfvo1ZE8z5F",
		URL:    "https://api.giphy.com/v1/gifs/random",
	}

	gif, _ := cl.FetchGIF(searchTerm)

	return gif
}
