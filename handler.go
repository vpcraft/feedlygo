package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handlerReadinessErr(w http.ResponseWriter, r *http.Request) {
	type ErrResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, http.StatusBadRequest, ErrResponse{
		Error: "Something went wrong."},
	)
}
