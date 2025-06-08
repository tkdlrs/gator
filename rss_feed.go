package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Create a client to make HTTP requests
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	// Make a GET request for the data. Provide context.
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	// add some headers....
	req.Header.Set("User-Agent", "gator")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// Read the data! So that it can be acutally used later.
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Set a place to hold the data.
	var rssFeed RSSFeed
	// unmarshal the data into an RSS Feed
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return nil, err
	}
	// Then 'un' escape the html/xml entities...
	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeed.Channel.Item[i] = item
	}
	//
	return &rssFeed, nil
}
