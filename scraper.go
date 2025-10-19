package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/vpcraft/feedlygo/internal/db"
)

func startScraping(
	db *db.Queries,
	concurrency int,
	fetchInterval time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrency, fetchInterval)
	ticker := time.NewTicker(fetchInterval)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)

		if err != nil {
			log.Println("error while fetching feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(wg, db, feed)
		}

		wg.Wait()
	}
}

func scrapeFeed(wg *sync.WaitGroup, db *db.Queries, feed db.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("error marking feed as fetched:", err)
		return
	}

	rssFeed, err := serializeFeedXML(feed.Url)
	if err != nil {
		log.Println("error serializing feed:", err)
		return
	}

	// Actually we should save on db, but let's test it first
	for _, item := range rssFeed.Channel.Item {
		log.Println("Found Post:", item.Title, "on feed", feed.Name)
	}
	log.Printf("Feed %s collected, %v posts found.", feed.Name, len(rssFeed.Channel.Item))
}
