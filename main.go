package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

var certPath = flag.String("cert", "config/cert.pem", "cert path")
var keyPath = flag.String("key", "config/key.pem", "key path")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("got request. remote addr: %s, proto: %s, headers: %+v", req.RemoteAddr, req.Proto, req.Header)
		io.WriteString(w, "{}")
	})

	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	log.Printf("listening at 8443")
	err := http.ListenAndServeTLS(":8443", *certPath, *keyPath, nil)
	log.Fatal(err)
}
