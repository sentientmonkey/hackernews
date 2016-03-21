package news

import (
	"testing"
	"time"
)

var (
	story        *Story
	displayStory *DisplayStory
)

func init() {
	story = &Story{
		Comments: []int{1, 2, 3},
		Time:     Timestamp(time.Date(2016, 3, 20, 16, 30, 0, 0, time.UTC)),
	}
	displayStory = story.ToDisplayStory(1)
}

func TestDisplayStoryCommentComent(t *testing.T) {
	expected := 3
	actual := displayStory.CommentCount()
	if actual != expected {
		t.Errorf("Expected CommentCount to be %d, got %d", expected, actual)
	}
}

func TestDisplayStoryFormattedTime(t *testing.T) {
	expected := "Sun Mar 20 16:30:00 2016"
	actual := displayStory.FormattedTime()
	if actual != expected {
		t.Errorf("Expected FormattedTime to be %q, got %q", expected, actual)
	}
}
