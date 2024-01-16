package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title        string        `xml:"title"`
		Link         string        `xml:"link"`
		Description  string        `xml:"description"`
		Language     string        `xml:"language"`
		RSSFeedItems []RSSFeedItem `xml:"item"`
	} `xml:"channel"`
}

type RSSFeedItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	response, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, err
	}
	defer response.Body.Close()
	dat, err := io.ReadAll(response.Body)
	if err != nil {
		return RSSFeed{}, err
	}
	rssFeed := RSSFeed{}
	xml.Unmarshal(dat, &rssFeed)
	return rssFeed, nil
}
