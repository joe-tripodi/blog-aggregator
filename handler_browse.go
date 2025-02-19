package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/joe-tripodi/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		var err error
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("failed to convert %s to a number", cmd.Args[0])
		}
	}

	browseParams := database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}

	posts, err := s.db.GetPostsForUser(context.Background(), browseParams)
	if err != nil {
		return fmt.Errorf("failed to get posts for user: %w:", err)
	}
	for _, post := range posts {
		if post.Title == "" {
			continue
		}
		fmt.Println("============================")
		printPost(post)
		fmt.Println("============================")
	}
	return nil
}

func printPost(post database.Post) {
	fmt.Println(post.Title)
	fmt.Println("Published:", post.PublishedAt)
	fmt.Println(post.Description)
}
