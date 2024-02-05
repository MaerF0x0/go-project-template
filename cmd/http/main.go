package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"

	responses "github.com/maerf0x0/go-project-template/internal/http"

	log "github.com/sirupsen/logrus"
)

const (
	// set version (or CommitHash, BuildTimestamp) at build time with
	// go build -ldflags="-X 'main.Version=0.0.0.rc123'"
	Version        = ""
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
	Port           = "3333" // as a string to save conversions
)

func logEntry() *log.Entry {
	return log.WithFields(log.Fields{
		"Version":    Version,
		"CommitHash": CommitHash,
	})
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	logEntry().Info("Example HTTP server")

	mux := http.NewServeMux()
	mux.HandleFunc("/health", sometimesFails)
	fmt.Println("Server is running on port: " + Port)
	err := http.ListenAndServe(":"+Port, mux)
	if err != nil {
		logEntry().WithError(err).Error("http server failed")
		return
	}

	logEntry().Info("http server stopped")

}

var (
	ErrRandFail = errors.New("random failure")
)

func sometimesFails(w http.ResponseWriter, r *http.Request) {
	jsonEnc := json.NewEncoder(w)
	if rand.Intn(100) < 10 {
		w.WriteHeader(http.StatusInternalServerError)
		err := jsonEnc.Encode(responses.NewError(ErrRandFail, http.StatusInternalServerError))
		if err != nil {
			logEntry().WithError(err).Error("Failed to encode error response")
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err := jsonEnc.Encode(responses.NewSuccess(http.StatusOK))
	if err != nil {
		logEntry().WithError(err).Error("Failed to encode success response")
	}
}
