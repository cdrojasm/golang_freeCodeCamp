package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cdrojasm/golang_freeCodeCamp/project/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("it was not possible to decode JSON request payload: %s", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create feed: %s", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feed, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Couldn't retrieve feed list due: %s", err))
		return
	}
	responseWithJSON(w, 200, sliceDatabaseFeedToSliceFeed(feed))
}
