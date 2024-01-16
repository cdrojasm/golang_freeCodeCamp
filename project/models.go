package main

import (
	"time"

	"github.com/cdrojasm/golang_freeCodeCamp/project/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"string"`
	APIKey    string    `json:"api_key"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func sliceDatabaseFeedToSliceFeed(dbFeeds []database.Feed) []Feed {
	var sliceFeed = make([]Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		sliceFeed[i] = Feed{
			ID:        dbFeed.ID,
			CreatedAt: dbFeed.CreatedAt,
			UpdatedAt: dbFeed.UpdatedAt,
			Name:      dbFeed.Name,
			Url:       dbFeed.Url,
			UserID:    dbFeed.UserID,
		}
	}
	return sliceFeed
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func feedFollowDatabaseToFeedFollow(dbFeedFollow database.FeedsFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.FeedID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		UserID:    dbFeedFollow.UserID,
		FeedID:    dbFeedFollow.FeedID,
	}
}

func sliceDatabaseFeedFollowsToSliceFeedFollows(dbFeedFollow []database.FeedsFollow) []FeedFollow {
	var sliceFeedFollows = make([]FeedFollow, len(dbFeedFollow))
	for i, dbFeedFollow := range dbFeedFollow {
		sliceFeedFollows[i] = FeedFollow{
			ID:        dbFeedFollow.ID,
			CreatedAt: dbFeedFollow.CreatedAt,
			UpdatedAt: dbFeedFollow.UpdatedAt,
			UserID:    dbFeedFollow.UserID,
			FeedID:    dbFeedFollow.FeedID,
		}
	}
	return sliceFeedFollows
}

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	Url         string    `json:"url"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func databasePostToPost(dbPost database.Post) Post {
	var description *string
	if dbPost.Description.Valid {
		description = &dbPost.Description.String
	}
	return Post{
		ID:          dbPost.FeedID,
		CreatedAt:   dbPost.CreatedAt,
		UpdatedAt:   dbPost.UpdatedAt,
		Title:       dbPost.Title,
		Description: description,
		PublishedAt: dbPost.PublishedAt,
		Url:         dbPost.Url,
		FeedID:      dbPost.FeedID,
	}
}

func sliceDatabasePostToSlicePost(dbPosts []database.Post) []Post {
	var returnPostSlice = make([]Post, len(dbPosts))
	for i, dbPost := range dbPosts {
		returnPostSlice[i] = databasePostToPost(dbPost)
	}
	return returnPostSlice
}
