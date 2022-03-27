package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	greetings := chi.URLParam(r, "greetings")
	rw.Write([]byte(greetings + " " + name))
}

func main() {
	r := chi.NewRouter()
	r.Get("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("chi"))
	})

	r.Get("/api/{greetings}", HelloWorld)

	log.Fatal(http.ListenAndServe(":8080", r))
}
