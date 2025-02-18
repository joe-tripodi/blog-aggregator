package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.Args) != 1 {
	// 	return fmt.Errorf("usage: %v <name>", cmd.Name)
	// }

	feedUrl := "https://www.wagslane.dev/index.xml" //cmd.Args[0]
	rssFeed, err := fetchFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("Unable to fetch feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", rssFeed)
	return nil
}
