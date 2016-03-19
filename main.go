package main

import (
	"flag"
	"fmt"
	"hackernews/news"
)

var (
	numberOfArticles int
	urlsOnly         bool
)

const (
	timeFormat = "Mon Jan 2 15:04:05 2006"
)

func init() {
	flag.IntVar(&numberOfArticles, "n", 10, "number of articles")
	flag.BoolVar(&urlsOnly, "u", false, "only show urls")
}

func main() {
	flag.Parse()

	showStories()
}

func showStories() {
	ids, err := news.TopStories()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	for i, id := range ids {
		story, err := news.GetStory(id)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else if urlsOnly {
			fmt.Printf("%s\n", story.Url)
		} else {
			fmt.Printf("%2d. %s (%s)\n", i+1, story.Title, story.Url)
			fmt.Printf("    %d points by %s %s | %d comments\n", story.Score, story.By, story.Time.Format(timeFormat), len(story.Comments))
		}
		if i+1 >= numberOfArticles {
			break
		}
	}
}
