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
	gif := getCarGIF()

	carMsg := `<h1>🚗 Car Valuation Service 🏎</h1>
<div>
<iframe src="%s" frameborder="0" width="800"></iframe>
</div>
<hr \>
<p>🤑 Based on our calculations, this is what your %s is worth: £%d 🎉</p>

<p>📱 Running on the following commit hash: %s.</p>


`
	carMsgB := []byte(fmt.Sprintf(carMsg, gif, model, value, commitHash))
	_, _ = w.Write(carMsgB)
}

func getCarGIF() string {
	cl := GiphyClient{
		APIKey: "UwelgUbg9VUspxHAf3E8XXfvo1ZE8z5F",
		URL:    "https://api.giphy.com/v1/gifs/random",
	}

	gif, _ := cl.FetchGIF("car")

	return gif
}
