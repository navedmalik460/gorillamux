package main

import (
	"fmt"
	"net/http"

	"gorillamux/handlerfunc"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")
	// Create a new Gorilla Mux router
	r := advancedRoutes()
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

// loggingMiddleware is a middleware function that logs the request method, URL path, and any variables.
// It takes an http.Handler as an argument and returns a new http.Handler.
// The middleware function wraps the original handler and adds logging functionality.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		muxVars := mux.Vars(r)
		fmt.Printf("Request: %s %s %s\n", r.Method, r.URL.Path, muxVars)
		next.ServeHTTP(w, r)
	})
}

// advance routes with handlers used to define own handlers
func advancedRoutes() *mux.Router {
	r := mux.NewRouter()
	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/getuser/{id}", handlerfunc.HandleGetuser).Methods("GET")

	r.HandleFunc("/hello/{name}/age/{age}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := vars["age"]
		fmt.Fprintf(w, "Hello, %s! You are %s years old.", name, age)
	}).Methods("GET")

	return r

}

// searchHandler to search query param from the URL.
func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	fmt.Fprintf(w, "Search query: %s", query)
}
