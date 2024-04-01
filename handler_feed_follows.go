package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bk167465/rss/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not create feed: %s", err))
		return
	}

	responseWithJson(w, 200, databaseFeedFollowToFeedFollow(feed))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
    feedFollowIDStr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIDStr)
	if err !=nil {
		responseWithError(w, 400, fmt.Sprintf("Could not parse feed folow ID: %s", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID: feedFollowId,
		UserID: user.ID,
	})
	if err !=nil {
		responseWithError(w, 400, fmt.Sprintf("Could not delete feed follow: %s", err))
		return
	}
	responseWithJson(w, 200, struct{}{})
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Could not get feed follows: %s", err))
		return
	}

	responseWithJson(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}
