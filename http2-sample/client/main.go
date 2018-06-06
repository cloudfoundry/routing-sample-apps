package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"golang.org/x/net/http2"
)

var host = flag.String("host", "", "Host addr")
var routerAddr = flag.String("routerAddr", "", "router addr")

func main() {
	flag.Parse()
	if *host == "" {
		log.Fatal("host must be provided")
	}
	if *routerAddr == "" {
		*routerAddr = *host
	}
	fmt.Println("host is set to: ", *host)
	fmt.Println("routerAddr is set to: ", *routerAddr)
	buf := &bytes.Buffer{}
	client := http.Client{
		// Skip TLS dial
		Transport: &http2.Transport{
			AllowHTTP: true,
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				conn, err := net.Dial(network, addr)
				return &cn{conn, buf}, err
			},
		},
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s", *routerAddr), nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Host = *host
	req.Proto = "HTTP/2"
	fmt.Printf("%+v\n", req)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(fmt.Errorf("error making request: %v", err))
	}

	defer resp.Body.Close()

	// read http2 frames first
	framer := http2.NewFramer(ioutil.Discard, buf)
	for {
		f, err := framer.ReadFrame()
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		switch err.(type) {
		case nil:
			log.Println(f)
		case http2.ConnectionError:
			// Ignore. There will be many errors of type "PROTOCOL_ERROR, DATA
			// frame with stream ID 0". Presumably we are abusing the framer.
		default:
			log.Println(err, framer.ErrorDetail())
		}
	}
}

type cn struct {
	net.Conn
	T io.Writer // receives everything that is read from Conn
}

func (w *cn) Read(b []byte) (n int, err error) {
	n, err = w.Conn.Read(b)
	w.T.Write(b)
	return
}
