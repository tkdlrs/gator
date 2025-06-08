package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.Args) != 1 {
	// 	return fmt.Errorf("usage: %s <url>", cmd.Name)
	// }
	// feed_url := cmd.Args[0]
	feed_url := "https://www.wagslane.dev/index.xml"
	//
	feed, err := fetchFeed(context.Background(), feed_url)
	if err != nil {
		return fmt.Errorf("error fetching RSS feed: %w", err)
	}
	//
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
