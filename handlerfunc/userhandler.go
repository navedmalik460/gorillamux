package handlerfunc

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// handleGetuser is a handler function that handles GET requests to the "/user/{id}" endpoint.
// It retrieves the user ID from the URL parameters and returns user data as JSON.
func HandleGetuser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// Simulate getting user data from a database
	user := map[string]string{
		"id":   id,
		"name": "John Doe",
	}

	// Return the user data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
