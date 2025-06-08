package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func handlerFeed(s *state, cmd command) error {
	// if len(cmd.Args) != 1 {
	// 	return fmt.Errorf("usage: %s <url>", cmd.Name)
	// }
	// feed_url := cmd.Args[0]
	feed_url := "https://www.wagslane.dev/index.xml"
	//
	fullRSSFeedData, err := fetchFeed(context.Background(), feed_url)
	if err != nil {
		return fmt.Errorf("error when fetching RSS feed: %v", err)
	}
	//
	fmt.Printf("%v and type of %T", fullRSSFeedData, fullRSSFeedData)
	fmt.Println()
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Set a place to hold the data.
	rssFeed := RSSFeed{}
	// get the data
	resp, err := http.Get(feedURL)
	if err != nil {
		return &rssFeed, err
	}
	defer resp.Body.Close()
	//
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return &rssFeed, err
	}
	// unmarshal the data into an RSS Feed
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return &rssFeed, err

	}
	//
	return &rssFeed, nil
}
