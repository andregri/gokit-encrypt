package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andregri/gokit-encrypt/helpers"
	httptransport "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
)

func main() {
	//
	logger := kitlog.NewLogfmtLogger(os.Stderr)

	// Declare service instance
	crypto_svc := helpers.EncryptServiceInstance{}
	svc := helpers.LogginMiddleware{Logger: logger, Next: crypto_svc}

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
