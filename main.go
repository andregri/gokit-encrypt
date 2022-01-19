package main

import (
	"log"
	"net/http"

	"github.com/andregri/gokit-encrypt/helpers"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	// Declare service instance
	svc := helpers.EncryptServiceInstance{}

	// Connect json decoder/encoder for request/response to encrypt endpoint
	encryptHandler := httptransport.NewServer(
		helpers.MakeEncryptEndpoint(svc),
		helpers.DecodeEncryptRequest,
		helpers.EncodeResponse,
	)

	// Connect json decoder/encoder for request/response to decrypt endpoint
	decryptHandler := httptransport.NewServer(
		helpers.MakeDecryptEndpoint(svc),
		helpers.DecodeDecryptRequest,
		helpers.EncodeResponse,
	)

	// Route incoming traffic
	http.Handle("/encrypt", encryptHandler)
	http.Handle("/decrypt", decryptHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
