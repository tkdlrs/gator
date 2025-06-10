package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tkdlrs/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	// Get the ID of the current user.
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	//
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed name> <url>", cmd.Name)
	}
	feedName := cmd.Args[0]
	url := cmd.Args[1]
	//
	feedArgs := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      feedName,
		Url:       url,
	}
	feed, err := s.db.CreateFeed(context.Background(), feedArgs)
	if err != nil {
		return fmt.Errorf("couldn't create a feed: %w", err)
	}
	//
	fmt.Println("Feed created successfully:")
	printFeed(feed)
	fmt.Println()
	fmt.Println("==================================")
	//
	return nil

}

func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
