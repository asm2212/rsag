package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/asm2212/rsag/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("scraping on %v goroutines every %s duration\n",
		concurrency,
		timeBetweenRequest)

	ticker := time.NewTicker(timeBetweenRequest)
	defer ticker.Stop()

	for {
		<-ticker.C

		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("Error fetching feeds:", err)
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

func scrapeFeed(
	wg *sync.WaitGroup,
	db *database.Queries,
	feed database.Feed,
) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error marking feed %d as fetched: %v\n", feed.ID, err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error fetching RSS feed %s: %v\n", feed.Url, err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("found post:", item.Title, "on feed:", feed.Name)
	}
	log.Printf("Feed %s collected, %v posts found\n", feed.Name, len(rssFeed.Channel.Item))
}
