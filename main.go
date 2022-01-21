package main

import (
	"log"
	"net/http"
	"os"

	"github.com/andregri/gokit-encrypt/helpers"
	httptransport "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"

	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Make new logger
	logger := kitlog.NewLogfmtLogger(os.Stderr)

	// Make new prometheus instrumentor
	fieldKeys := []string{"method", "error"}

	requestCount := kitprometheus.NewCounterFrom(prometheus.CounterOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)

	requestLatency := kitprometheus.NewSummaryFrom(prometheus.SummaryOpts{
		Namespace: "encryption",
		Subsystem: "my_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)

	// Declare service instance
	crypto_svc := helpers.EncryptServiceInstance{}
	logging_svc := helpers.LogginMiddleware{Logger: logger, Next: crypto_svc}
	svc := helpers.InstrumentationMiddleware{
		RequestCount:   requestCount,
		RequestLatency: requestLatency,
		Next:           logging_svc,
	}

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
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
