package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Greeting: Welcome {%s}!\n%s", vars["name"], time.Now()。Format("2006-01-02 15:04:05"))
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
