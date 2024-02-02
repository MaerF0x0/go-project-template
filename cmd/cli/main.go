package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.WithFields(log.Fields{
		"Args": os.Args,
	}).Info("Example CLI")
}
