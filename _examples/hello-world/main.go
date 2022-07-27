package main

import (
	"net/http"

	phi "github.com/PhilipJovanovic/phi"
	"github.com/PhilipJovanovic/phi/middleware"
)

func main() {
	r := phi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":3333", r)
}
