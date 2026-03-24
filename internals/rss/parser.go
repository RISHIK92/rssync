package rss

import (
	"encoding/xml"
	"fmt"
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

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	lastUpdated := state.LastUpdated()

	var rss RSS
	arr := []Item{}

	err = xml.Unmarshal(bodyBytes, &rss)

	if err != nil {
		return nil, err
	}

	if len(rss.Channel.Items) == 0 {
		return arr, nil
	}

	for i:=range rss.Channel.Items {
		if formatter(rss.Channel.Items[i].PubDate).Before(formatter(lastUpdated)) {
			break
		}
		arr = append(arr, rss.Channel.Items[i])
	}

	loc, err := time.LoadLocation("GMT")
	if err != nil {
		fmt.Println("Error loading timezone:", err)
	}

	nowGMT := time.Now().In(loc)

	state.UpdateLastUpdated(nowGMT.Format(time.RFC1123))

	return arr, nil
}