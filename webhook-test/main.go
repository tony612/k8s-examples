package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

var certPath = flag.String("cert", "config/cert.pem", "cert path")
var keyPath = flag.String("key", "config/key.pem", "key path")

type AdmissionReview struct {
	ApiVersion string   `json:"apiVersion,omitempty"`
	Kind       string   `json:"kind"`
	Request    Request  `json:"request,omitempty"`
	Response   Response `json:"response,omitempty"`
}

type Request struct {
	Uid string `json:"uid"`
}

type Response struct {
	Uid     string `json:"uid"`
	Allowed bool   `json:"allowed"`
}

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("got request. remote addr: %s, proto: %s, headers: %+v", req.RemoteAddr, req.Proto, req.Header)
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("fail to read req body: %s", err)
		}
		if os.Getenv("PRINT_REQ") != "" {
			log.Printf("got body: %s", string(body))
		}

		r := AdmissionReview{}
		err = json.Unmarshal(body, &r)
		if err != nil {
			log.Printf("fail to unmarshal req body: %s", err)
		}

		log.Printf("uuid: %s", r.Request.Uid)

		resp := AdmissionReview{
			ApiVersion: r.ApiVersion,
			Kind:       r.Kind,
			Response: Response{
				Uid:     r.Request.Uid,
				Allowed: os.Getenv("NOT_ALLOWED") == "",
			},
		}
		if !resp.Response.Allowed {
			log.Printf("send disallowed")
		}
		err = json.NewEncoder(w).Encode(&resp)
		if err != nil {
			log.Printf("fail to marshal response: %s", err)
		}
	})

	// One can use generate_cert.go in crypto/tls to generate cert.pem and key.pem.
	log.Printf("listening at 8443")
	err := http.ListenAndServeTLS(":8443", *certPath, *keyPath, nil)
	log.Fatal(err)
}
