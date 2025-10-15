package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/vpcraft/feedlygo/internal/auth"
	"github.com/vpcraft/feedlygo/internal/db"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Fullname string `json:"fullname"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while parsing JSON: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		Fullname:  params.Fullname,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while creating user: %v", err))
		return
	}

	respondWithJSON(w, 201, serializerUser(user))
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
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

	if user == (db.User{}) {
		respondWithError(w, 400, "User not found")
		return
	}

	respondWithJSON(w, 200, serializerUser(user))
}
