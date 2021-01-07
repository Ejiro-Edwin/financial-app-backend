package main

import (
	"github.com/ejiro-edwin/Financial-App/internal/api"
	"github.com/ejiro-edwin/Financial-App/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

// create server object and start listener
func main() {

	logrus.SetLevel(logrus.DebugLevel)

	logrus.WithField("version", config.Version).Debug("Starting server.")

	router, err := api.NewRouter()
	if err != nil {

		logrus.WithError(err).Fatal("Error building router")
	}

	const addr = "0.0.0.0:8088"
	server := http.Server{
		Handler: router,
		Addr:    addr,
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logrus.WithError(err).Error("Server Listens")
	}
}
