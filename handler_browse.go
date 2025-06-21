package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tkdlrs/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}
	//
	getPostsArgs := database.GetUsersPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.db.GetUsersPosts(context.Background(), getPostsArgs)
	if err != nil {
		return fmt.Errorf("couldn't get users posts: %w", err)
	}
	//
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		printPost(post)
		fmt.Println("==================================")
	}
	//
	return nil
}

func printPost(post database.GetUsersPostsRow) {
	fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
	fmt.Printf("--- %s ---\n", post.Title)
	fmt.Printf("    %v \n", post.Description.String)
	fmt.Printf("Link: %s\n", post.Url)
}
