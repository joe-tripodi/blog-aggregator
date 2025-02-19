package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joe-tripodi/gator/internal/database"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("failed to get user for feed: %w", err)
		}

		fmt.Println("=====================================")
		printFeed(feed, user)
		fmt.Println()
		fmt.Println("=====================================")
	}
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	rssFeed, err := s.db.CreateFeed(context.Background(), feed)
	if err != nil {
		return fmt.Errorf("Unable to created feed: %w", err)
	}

	printFeed(rssFeed, user)

	// now we need to create a feed follow
	// I could do this by updating the query OR I do this here
	feedFollow := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    rssFeed.ID,
	}

	_, err = s.db.CreateFeedFollows(context.Background(), feedFollow)
	if err != nil {
		return fmt.Errorf("unable to create feed follow: %w", err)
	}

	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* LastFetched:   %v\n", feed.LastFetchedAt)
	fmt.Printf("* Username:      %s\n", user.Name)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
	fmt.Printf("* ID:            %s\n", feed.ID)
}
