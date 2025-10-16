package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/vpcraft/feedlygo/internal/db"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), db.CreateFeedParams{
		ID:     uuid.New(),
		Name:   params.Name,
		Url:    params.URL,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while creating feed: %v", err))
		return
	}

	respondWithJSON(w, 200, serializerFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeedByID(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		ID uuid.UUID `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.GetFeed(r.Context(), params.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while getting feed: %v", err))
		return
	}

	respondWithJSON(w, 200, serializerFeed(feed))
}

func (apiCfg *apiConfig) handlerGetAllFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetAllFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while getting feeds: %v", err))
		return
	}

	serializedFeeds := serializerFeeds(feeds)
	respondWithJSON(w, 200, serializedFeeds)
}
