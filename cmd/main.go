package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rishik92/rssync/internals/mailer"
	"github.com/rishik92/rssync/internals/rss"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found. Proceeding with existing environment variables.")
	}

	url := "https://arpitbhayani.me/rss.xml"
	result, err := rss.ParseRSSFeed(url)

	if len(result)==0 {
		fmt.Println("No new items found in the RSS feed.")
	}

	if err != nil {
		panic(err)
	}
	
	for i:=range result {
		item := result[i]
		success, err1 := mailer.SendEmail("rishik3555@gmail.com",[]string{"rishik3555@gmail.com"}, item.Title, item.Link, item.PubDate, item.Description)
		if err1 != nil {
			fmt.Printf("Error sending email for item %s: %v\n", item.Title, err1)
			return
		}
		fmt.Println("Email sent successfully for item:", item.Title, "Success:", success)
	}
}