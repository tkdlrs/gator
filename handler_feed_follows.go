package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tkdlrs/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	//
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	// Get the ID of the current Feed based on the URL
	url := cmd.Args[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}
	// create a new feed follow record for current user.
	feedFollowArgs := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollowsRow, err := s.db.CreateFeedFollow(context.Background(), feedFollowArgs)
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	// Print the name of the feed and current users
	fmt.Println("Feed follow created:")
	printFeedFollow(feedFollowsRow.UserName, feedFollowsRow.FeedName)
	//
	return nil
}

func handlerListFeedFollows(s *state, cmd command, user database.User) error {
	//
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}
	//
	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}
	//
	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}
	//
	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
