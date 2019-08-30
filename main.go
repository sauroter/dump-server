package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	fmt.Printf(string(dump))
	fmt.Println()
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
