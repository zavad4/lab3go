package forums

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zavad4/lab3go/tree/main/server/tools"
)

// Forums HTTP handler.
type HttpHandlerFunc http.HandlerFunc

// HttpHandler creates a new instance of forums HTTP handler.
func HttpHandler(store *Store) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListForums(store, rw)
		} else if r.Method == "POST" {
			handleUserCreate(r, rw, store)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleUserCreate(r *http.Request, rw http.ResponseWriter, store *Store) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Printf("Error decoding user input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := store.CreateUser(u.Name, u.Topics)
	if err == nil {
		tools.WriteJsonOk(rw, &u)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListForums(store *Store, rw http.ResponseWriter) {
	res, err := store.ListForums()
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
