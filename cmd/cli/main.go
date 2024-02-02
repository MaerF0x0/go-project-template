package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	// set version at build time with
	// go build -ldflags="-X 'main.Version=0.0.0.rc123'"
	Version        string
	CommitHash     = "n/a"
	BuildTimestamp = "n/a"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	log.WithFields(log.Fields{
		"Args":       os.Args,
		"Version":    Version,
		"CommitHash": CommitHash,
	}).Info("Example CLI")
}
