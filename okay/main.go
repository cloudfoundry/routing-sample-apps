package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", okay)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func okay(w http.ResponseWriter, req *http.Request) {
	log.Printf("responding to %s\n", req.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello\n")
}
