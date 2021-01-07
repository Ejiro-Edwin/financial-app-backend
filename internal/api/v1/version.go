package v1

import (
	"encoding/json"
	"github.com/ejiro-edwin/Financial-App/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

// API for returning  version , when server starts we set the version and then use it if necessary

var versionJSON []byte

type ServerVersion struct {
	Version string `json:"version"`
}

func init() {
	var err error
	versionJSON, err = json.Marshal(ServerVersion{Version: config.Version})

	if err != nill {
		panic(err)
	}
}

// versionHandler serves version information
func VersionHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	if _, err := w.Write(versionJSON); err != nil {
		logrus.WithError(err).Debug("Error writing version")
	}

}
