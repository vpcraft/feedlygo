package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/vpcraft/feedlygo/internal/db"
)

func (apiCfg *apiConfig) handlerFollowToFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while parsing JSON: %v", err))
		return
	}

	follow, err := apiCfg.DB.FollowToFeed(r.Context(), db.FollowToFeedParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while following feed: %v", err))
		return
	}

	respondWithJSON(w, 201, serializerFollow(follow))
}

func (apiCfg *apiConfig) handlerUnfollowFromFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while parsing JSON: %v", err))
		return
	}

	err = apiCfg.DB.UnfollowFromFeed(r.Context(), db.UnfollowFromFeedParams{
		UserID: user.ID,
		FeedID: params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while following feed: %v", err))
		return
	}

	respondWithJSON(w, 204, struct{}{})
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user db.User) {
	follows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error while getting follows: %v", err))
		return
	}

	respondWithJSON(w, 200, serializerFollows(follows))
}
