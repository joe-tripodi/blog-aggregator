package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/joe-tripodi/gator/internal/database"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("fetching feed: %v\n", nextFeed.Name)

	markFeedParams := database.MarkFeedFetchedParams{
		ID: nextFeed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}

	err = s.db.MarkFeedFetched(context.Background(), markFeedParams)
	if err != nil {
		return fmt.Errorf("unable to fetch feed: %w", err)
	}

	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("Feed: %+v\n", rssFeed.Channel.Title)
	for _, item := range rssFeed.Channel.Item {
		fmt.Println(item.Title)
	}
	fmt.Println("=======================================")

	return nil
}

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to parse time: %w", err)
	}

	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
