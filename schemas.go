package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/vpcraft/feedlygo/internal/db"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Fullname  string    `json:"fullname"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Follow struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func serializerUser(user db.User) User {
	return User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Fullname:  user.Fullname,
	}
}

func serializerFeed(feed db.Feed) Feed {
	return Feed{
		ID:        feed.ID,
		Name:      feed.Name,
		URL:       feed.Url,
		UserID:    feed.UserID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
	}
}

func serializerFeeds(feeds []db.Feed) []Feed {
	resFeeds := []Feed{}
	for _, dbFeed := range feeds {
		resFeeds = append(resFeeds, serializerFeed(dbFeed))
	}

	return resFeeds
}

func serializerFollow(follow db.FeedFollow) Follow {
	return Follow{
		ID:        follow.ID,
		UserID:    follow.UserID,
		FeedID:    follow.FeedID,
		CreatedAt: follow.CreatedAt,
		UpdatedAt: follow.UpdatedAt,
	}
}

func serializerFollows(follows []db.FeedFollow) []Follow {
	resFollows := []Follow{}
	for _, dbFollow := range follows {
		resFollows = append(resFollows, serializerFollow(dbFollow))
	}

	return resFollows
}
