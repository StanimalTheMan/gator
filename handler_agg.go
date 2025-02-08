package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/StanimalTheMan/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs> aka duration string, like 1s, 1m, 1h, etc", cmd.Name)
	}
	time_between_reqs := cmd.Args[0]
	parsed_time_between_reqs, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("error returning parsed time%w", err)
	}
	ticker := time.NewTicker(parsed_time_between_reqs)
	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Couldn't get next feeds to fetch", err)
		return
	}
	log.Println("Found a feed to fetch!")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
