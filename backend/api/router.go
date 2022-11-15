package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	s := r.PathPrefix("/api/v1").Subrouter()

	api := &API{}

	s.HandleFunc("/", api.HomeHandler)
	/*
		s.HandleFunc("/assets/{assetID}/spectra/{id}/data/{t}", api.SpecDataHandler).Name("specData")

		s.HandleFunc("/devices/{deviceID}/conf", api.ImportConfHandler).Name("importDeviceConf").Methods("POST")
	*/

	return s
}
