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

var port = "8080"

var version = "v0.0.2"

func main() {
	http.HandleFunc("/hello", helloWorldHandler)
	http.HandleFunc("/hello-json", helloWorldJsonHandler)
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health/live", healthHandler)
	http.HandleFunc("/health/ready", healthHandler)
	http.HandleFunc("/", homepageHandler)

	if os.Getenv("JSON_LOG") == "" {
		// disable json logging
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	log.Info().Msgf("Running on %v architecture, app version %v", runtime.GOARCH, version)
	log.Info().Msgf("Listening on port %v", port)
	log.Info().Msg("Listening on /hello /hello-json /metrics /health/live /health/ready")
	log.Info().Msg(`Set env variable "JSON_LOG=true" to enable json logging`)
	log.Info().Msgf("env SECRET value: %v", os.Getenv("SECRET"))
	err := http.ListenAndServe(":"+port, nil)
	log.Fatal().Err(err).Msg("Could not start server")
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("Received incoming web request %v on path %v", r.Method, r.URL.Path)
	_, _ = w.Write([]byte("Hello version: " + version))
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("Received incoming web request %v on path %v", r.Method, r.URL.Path)
	// list all endpoints
	_, _ = w.Write([]byte("Listening on /hello /hello-json /metrics /health/live /health/ready"))
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
	_, _ = w.Write(jsonData)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
