package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	port := flag.Int("port", 8080, "port to listen on")

	flag.Parse()

	s := http.NewServeMux()

	s.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, r.RemoteAddr)

		w.WriteHeader(http.StatusOK)
	})

	address := fmt.Sprintf(":%d", *port)

	http.ListenAndServe(address, s)
}
