package app

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func RunHello(cmdFlag string, port string) {
	log.Println("Botting up hello app...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		hello(w, cmdFlag)
	})

	log.Printf("Listening on port %s\n", port)
	_ = http.ListenAndServe(port, r)
}

func hello(w http.ResponseWriter, cmdFlag string) {
	commitHash := os.Getenv("COMMIT_HASH")
	name := os.Getenv("name")
	if name == "" {
		name = "Cuvva"
	}
	helloMsg := `<h1>Hello %s! 🙌</h1>
<hr \>
<p>🎉 This service was invoked from main with the %s command.</p>

<p>📱 This is running on the following commit hash: %s.</p>
`
	helloMsgB := []byte(fmt.Sprintf(helloMsg, name, cmdFlag, commitHash))
	_, _ = w.Write(helloMsgB)
}
