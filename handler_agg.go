package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/tkdlrs/gator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) < 1 || len(cmd.Args) > 2 {
		return fmt.Errorf("usage: %s <time_between_reqs>", cmd.Name)
	}
	//
	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}
	log.Printf("Collecting feeds every %s...", timeBetweenRequests)
	//
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
func scrapeFeeds(s *state) {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		log.Println("Could not fetch next feeds", err)
		return
	}
	log.Println("Found a feed to fetch")
	scrapeFeed(s.db, feed)
}

func scrapeFeed(db *database.Queries, feed database.Feed) {
	_, err := db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}
	//
	feedData, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		log.Printf("Could not collect feed %s: %v", feed.Name, err)
		return
	}
	for _, item := range feedData.Channel.Item {
		fmt.Printf("Found post: %s\n", item.Title)
		// Save it to the 'posts' table
		// Fix description
		cleanDescription := sql.NullString{
			String: item.Description,
			Valid:  true,
		}
		// Fix publish date. It needs to be cleaned/made consistent
		timeLayout := "Mon, 02 Jan 2006 15:04:05 -0700"
		cleanTime, err := time.Parse(timeLayout, item.PubDate)
		if err != nil {
			fmt.Println("")
			fmt.Println(item.PubDate, "item.PubDate")
			fmt.Println("")

			log.Printf("Could not clean time: %s", err)
		}

		// stuff needed for the post
		postArgs := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: cleanDescription,
			PublishedAt: cleanTime,
			FeedID:      feed.ID,
		}
		post, err := db.CreatePost(context.Background(), postArgs)
		if err != nil {
			log.Printf("could not create post: %v\n", err)
		}
		fmt.Printf("post: %s saved to posts table\n", post.Url)

	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(feedData.Channel.Item))
}
