package main

import (
	"fmt"
	"net/http"

	"github.com/vpcraft/feedlygo/internal/auth"
	"github.com/vpcraft/feedlygo/internal/db"
)

type authedHandler func(http.ResponseWriter, *http.Request, db.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth_key, err := auth.GetBasicAuthAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Error while getting API key: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), auth_key)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Error while getting user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
