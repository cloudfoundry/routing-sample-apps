package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", closer)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func closer(w http.ResponseWriter, req *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	conn, _, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusTeapot)
		return
	}

	log.Printf("closing connection from %s\n", req.RemoteAddr)
	conn.Close()
}
