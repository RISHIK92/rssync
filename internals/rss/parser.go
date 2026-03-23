package rss

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/rishik92/rssync/internals/state"
)

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type RSS struct {
	Channel Channel `xml:"channel"`
}

func formatter(date string) (time.Time) {
	parsedDate, err := time.Parse(time.RFC1123, date)
	if err != nil {
		return time.Time{}
	}
	return parsedDate
}

func ParseRSSFeed(url string) ([]Item, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err1 := io.ReadAll(response.Body)

	if err1 != nil {
		return nil, err1
	}

	lastUpdated := state.LastUpdated()

	var rss RSS
	arr := []Item{}

	xml.Unmarshal(bodyBytes, &rss)

	for i := 0; i < len(rss.Channel.Items); i++ {
		if formatter(rss.Channel.Items[i].PubDate).Before(formatter(lastUpdated)) {
			break
		}
		arr = append(arr, rss.Channel.Items[i])
	}
	state.UpdateLastUpdated(rss.Channel.Items[0].PubDate)

	return arr, nil
}