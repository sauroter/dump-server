package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
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

	if len(os.Args) < 2 {
		log.Fatal("port required")
		os.Exit(1)
	}
	arg := os.Args[1]
	port := -1
	fmt.Sscanf(arg, "%d", &port)
	if port <= 0 {
		log.Fatal("port number must be in range")
		os.Exit(1)
	}

	log.Printf("started dump server on port %d\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
