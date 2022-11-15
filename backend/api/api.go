package api

import (
	"fmt"
	"net/http"
)

type API struct {
}

func (api *API) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the API basic handler!")
}
