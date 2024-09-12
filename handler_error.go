package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 400, "Sometthing went wrong :(")
}
