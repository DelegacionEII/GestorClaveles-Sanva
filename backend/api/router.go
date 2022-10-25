package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	api := &API{}

	r.HandleFunc("/", api.HomeHandler)
	/*
		s.HandleFunc("/assets/{assetID}/spectra/{id}/data/{t}", api.SpecDataHandler).Name("specData")

		s.HandleFunc("/devices/{deviceID}/conf", api.ImportConfHandler).Name("importDeviceConf").Methods("POST")
	*/

	http.Handle("/", r)
	return r
}
