package api

import (
	v1 "github.com/ejiro-edwin/Financial-App/internal/api/v1"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() (http.Handler, error) {
	router := mux.NewRouter()
	router.HandleFunc("/version", v1.VersionHandler)
	return router, nil

}
