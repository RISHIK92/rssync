package state

import (
	"fmt"
	"os"
)

type State struct {
	Date 		string
}

func LastUpdated() string {
	date, _ := os.ReadFile("last_updated.txt")
	return string(date)
}

func UpdateLastUpdated(date string) {
	err := os.WriteFile("last_updated.txt", []byte(date), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}