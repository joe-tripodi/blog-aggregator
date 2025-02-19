package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/joe-tripodi/gator/internal/database"
)

func handlerFollowing(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to get current user: %w", err)
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("failed to get follows for user: %w", err)
	}
	fmt.Println(following)

	for _, follow := range following {
		printFollow(follow)
	}
	return nil
}

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("unable to get feed by id: %w", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("unable to get the current user by name: %w", err)
	}

	feedFollow := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	dbFeedFollow, err := s.db.CreateFeedFollows(context.Background(), feedFollow)
	if err != nil {
		return fmt.Errorf("failed to create feed follow: %w", err)
	}

	fmt.Println("feed follow created")
	printFeedFollow(dbFeedFollow)

	return nil
}

func printFeedFollow(feedFollow database.CreateFeedFollowsRow) {
	fmt.Printf("* ID:            %s\n", feedFollow.ID)
	fmt.Printf("* Created:       %v\n", feedFollow.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feedFollow.UpdatedAt)
	fmt.Printf("* UserID:        %s\n", feedFollow.UserID)
	fmt.Printf("* Username:      %s\n", feedFollow.UserName)
	fmt.Printf("* FeedID:        %s\n", feedFollow.FeedID)
	fmt.Printf("* Feed:          %s\n", feedFollow.FeedName)
}

func printFollow(feedFollow database.GetFeedFollowsForUserRow) {
	fmt.Printf("* ID:            %s\n", feedFollow.ID)
	fmt.Printf("* Created:       %v\n", feedFollow.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feedFollow.UpdatedAt)
	fmt.Printf("* UserID:        %s\n", feedFollow.UserID)
	fmt.Printf("* Username:      %s\n", feedFollow.UserName)
	fmt.Printf("* FeedID:        %s\n", feedFollow.FeedID)
	fmt.Printf("* Feed:          %s\n", feedFollow.FeedName)
}
