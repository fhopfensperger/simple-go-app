package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"

	"github.com/rs/zerolog/log"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const port = "8080"

var version = "v0.0.2"

func main() {
	http.HandleFunc("/hello", helloWorldHandler)
	http.HandleFunc("/hello-json", helloWorldJsonHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health/live", healthHandler)
	http.HandleFunc("/health/ready", healthHandler)

	if os.Getenv("JSON_LOG") == "" {
		// disable json logging
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	log.Info().Msgf("Running on %v architecture, app version %v", runtime.GOARCH, version)
	log.Info().Msgf("Listening on port %v", port)
	log.Info().Msg("Listening on /hello /hello-json /metrics /health/live /health/ready")
	log.Info().Msg(`Set env variable "JSON_LOG=true" to enable json logging`)
	log.Info().Msgf("env SECRET value: %v", os.Getenv("SECRET"))
	http.ListenAndServe(":"+port, nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("Received incoming web request %v on path %v", r.Method, r.URL.Path)
	w.Write([]byte("Hello version: " + version))
}

func helloWorldJsonHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("Received incoming web request %v on path %v", r.Method, r.URL.Path)
	data := map[string]string{
		"message": "Hello",
		"version": version,
	}

	jsonData, err := json.Marshal(&data)
	if err != nil {
		log.Err(err).Msg("")
	}
	w.Write(jsonData)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
