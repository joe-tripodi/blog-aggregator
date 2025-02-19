package main

import (
	"context"
	"fmt"

	"github.com/joe-tripodi/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	url := cmd.Args[0]
	err := s.db.Unfollow(context.Background(), database.UnfollowParams{
		UserID: user.ID,
		Url:    url,
	})

	if err != nil {
		return fmt.Errorf("unable to unfollow feed: %w", err)
	}

	return nil
}
