package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RunGiffer(port string) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderGif(w, r)
	})

	log.Printf("Listening on port %s", port)
	_ = http.ListenAndServe(port, r)
}

func renderGif(w http.ResponseWriter, r *http.Request) {
	gif := getRandomGIF()

	memeMsg := `<h1>
<iframe src="%s" frameborder="0" width="%d" height="%d"></iframe>
</h1>
`
	memeMsgB := []byte(fmt.Sprintf(memeMsg, gif, 1000, 1000))
	_, _ = w.Write(memeMsgB)
}

func getRandomGIF() string {
	cl := GiphyClient{
		APIKey: "UwelgUbg9VUspxHAf3E8XXfvo1ZE8z5F",
		URL:    "https://api.giphy.com/v1/gifs/random",
	}

	gif, _ := cl.FetchGIF("meme")

	return gif
}
