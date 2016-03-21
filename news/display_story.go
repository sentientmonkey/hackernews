package news

// Wraps a Story with additional information (like story number)
// As well as provide helper functions for formatting to a template.
type DisplayStory struct {
	*Story
	Number int
}

func (story *Story) ToDisplayStory(number int) *DisplayStory {
	return &DisplayStory{
		Number: number,
		Story:  story,
	}
}

func (story *DisplayStory) CommentCount() int {
	return len(story.Comments)
}

func (story *DisplayStory) FormattedTime() string {
	return story.Time.Format("Mon Jan 2 15:04:05 2006")
}
