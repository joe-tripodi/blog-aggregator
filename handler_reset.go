package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to reset db: %w", err)
	}
	fmt.Println("successfully reset the db")

	feedUrl := cmd.Args[0]
	rssFeed, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Unable to fetch feed: %w", err)
	}
	fmt.Println(rssFeed)
	fmt.Println(*rssFeed)
	return nil
}
