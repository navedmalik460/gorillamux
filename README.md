Introduction
Gorilla Mux is a powerful and flexible router for the Go programming language. It provides advanced routing capabilities such as URL parameters, query parameters, middleware support, request handling, subrouters, and more. It is widely used in enterprise-level applications due to its efficiency and scalability.
Installation
To install Gorilla Mux, use the following command:
 go get -u github.com/gorilla/mux
Then import it in your Go project:
import "github.com/gorilla/mux"

Creating a Router
To create a new router instance, use:
package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	},
	).Methods("GET")
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Naved")
	},
	).Methods("GET")
  http.ListenAndServe(":8080", r)
}
Sumple gorilla mux hello api end point "http://localhost:8080/" and "http://localhost:8080/hello"

Handle advance route with it's own handle function ratherthan writing in same function and call to make code neet and clean
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
// check handlerfunc for the handling of user calls
this function has subroutes features as well to understand subroutes please follow this code 
userRouter := r.PathPrefix("/user").Subrouter()
userRouter.HandleFunc("/getuser/{id}", handlerfunc.HandleGetuser).Methods("GET")

