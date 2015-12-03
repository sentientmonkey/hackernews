package main

import (
	"flag"
	"fmt"
	"hackernews/news"
)

var numberOfArticles = flag.Int("n", 10, "number of articles")

func main() {
	flag.Parse()

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
			fmt.Printf("    %d points by %s %s | %d comments\n", story.Score, story.By, story.Time.Format("Mon Jan 2 15:04:05 2006"), len(story.Comments))
		}
		if i+1 >= *numberOfArticles {
			break
		}
	}
}
