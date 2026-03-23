package main

import (
	"fmt"

	"github.com/rishik92/rssync/internals/rss"
)

func main() {
	url := "https://arpitbhayani.me/rss.xml"
	result, _ := rss.ParseRSSFeed(url)
	fmt.Println(result)
}