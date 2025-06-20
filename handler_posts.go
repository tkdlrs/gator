package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/tkdlrs/gator/internal/database"
)

func handlerListPosts(s *state, cmd command, user database.User) error {
	limit := int32(2)
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s *<feed_limit> (optional)", cmd.Name)
	}
	if len(cmd.Args) == 1 {
		num, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("error converting string to int: %v", err)
		}
		limit = int32(num)
	}
	//
	getPostsArgs := database.GetUsersPostsParams{
		ID:    user.ID,
		Limit: limit,
	}
	posts, err := s.db.GetUsersPosts(context.Background(), getPostsArgs)
	if err != nil {
		return fmt.Errorf("couldn't list posts: %w", err)
	}
	//
	if len(posts) == 0 {
		fmt.Println("No posts found.")
		return nil
	}
	//
	fmt.Printf("Found post\n")
	for _, post := range posts {
		printPost(post)
		fmt.Println("==================================")
	}
	//
	return nil
}

func printPost(post database.GetUsersPostsRow) {
	fmt.Printf("* ID:           		%s\n", post.ID)
	fmt.Printf("* URL:           		%s\n", post.Url)
	fmt.Printf("* Description:          %v\n", post.Description)
}
