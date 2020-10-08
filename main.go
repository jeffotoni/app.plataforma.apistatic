package main

import (
	"log"
	"net/http"
	"strings"

	_ "github.com/jeffotoni/app.plataforma.apistatic/statik"
	"github.com/rakyll/statik/fs"
)

var (
	port = ":8080"
)

func main() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	fs := http.FileServer(statikFS)

	mux.HandleFunc("/ping", Ping)

	// mostra no browser localhost:8080/static
	mux.Handle("/", http.StripPrefix("/", DisabledFs(fs)))
	log.Println("Run Server:", port)
	http.ListenAndServe(port, mux)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`pong`))
}

func DisabledFs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
