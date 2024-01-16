package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cdrojasm/golang_freeCodeCamp/project/internal/database"
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
		responseWithError(w, 400, fmt.Sprintf("it was not possible to decode JSON request payload: %s", err))
		return
	}
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %s", err))
		return
	}

	responseWithJSON(w, 201, feedFollowDatabaseToFeedFollow(feedFollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollowsByID(w http.ResponseWriter, r *http.Request, user database.User) {
	feedsFollow, err := apiCfg.DB.GetFeedFollowsByUserID(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't get feed follow: %s", err))
		return
	}

	responseWithJSON(w, 200, sliceDatabaseFeedFollowsToSliceFeedFollows(feedsFollow))
}

func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	if feedFollowIDStr == "" {
		responseWithError(w, 400, "It was not provided a feed follow id to delete.")
		return
	}
	feedFollow, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("It was not possible to parse provided string to UUID due: %s", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollow,
		UserID: user.ID,
	})
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow: %s", err))
		return
	}

	responseWithJSON(w, 200, struct{}{})
}
