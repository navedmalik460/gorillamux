package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	},
	).Methods("GET")
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Naved")
	},
	).Methods("GET")
	r.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		fmt.Fprintf(w, "Hello, %s!", name)
	},
	).Methods("GET")

	http.ListenAndServe(":8080", r)

}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		muxVars := mux.Vars(r)
		fmt.Printf("Request: %s %s %s\n", r.Method, r.URL.Path, muxVars)
		next.ServeHTTP(w, r)
	})
}
