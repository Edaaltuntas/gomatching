package web

import "github.com/gorilla/mux"

// New Router to match routes to handlers
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/result", ResultHandler)
	r.HandleFunc("/api/signup", SignUpHandler)
	r.HandleFunc("/api/matching", MatchingHandler)
	return r
}
