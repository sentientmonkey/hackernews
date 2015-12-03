package main

import (
	"fmt"
	"hackernews/news"
)

func main() {
	ids, err := news.TopStories()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	for i, id := range ids {
		story, err := news.GetStory(id)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("%2d. %s (%s)\n", i+1, story.Title, story.Url)
			fmt.Printf("    %d points by %s %s | %d comments\n", story.Score, story.By, story.Time.String(), len(story.Comments))
		}
		if i+1 >= 10 {
			break
		}
	}
}
