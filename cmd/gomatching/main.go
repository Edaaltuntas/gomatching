package main

import (
	"log"
	"net/http"
	"os"

	"github.com/edaaltuntas/gomatching/pkg/web"
)

func main() {
	r := web.NewRouter()
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}
