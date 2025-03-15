package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("could not fetch RSS feed: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
