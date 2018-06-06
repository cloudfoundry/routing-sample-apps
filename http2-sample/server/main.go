package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hkwi/h2c"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// only HTTP2 reqs
		if r.ProtoMajor != 2 {
			log.Println("Not a HTTP/2 request, rejected!")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println(r)
		io.WriteString(w, "hello\n")
	})
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), &h2c.Server{}))
}
