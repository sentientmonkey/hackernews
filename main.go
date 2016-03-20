package main

import (
	"flag"
	"hackernews/news"
	"html/template"
	"log"
	"os"
)

var (
	numberOfArticles int
	urlsOnly         bool
)

func init() {
	flag.IntVar(&numberOfArticles, "n", 10, "number of articles")
	flag.BoolVar(&urlsOnly, "u", false, "only show urls")
}

var templateMap map[string]*template.Template

const storyTemplateContent = `{{.Number | printf "%2d"}}. {{.Title}} ({{.Url}})
    {{.Score}} points by {{.By}} {{.FormattedTime}} | {{.CommentCount}} comments
`
const urlTemplateContent = `{{.Url}}
`

func main() {
	flag.Parse()

	if err := registerTemplates(); err != nil {
		log.Println(err)
		return
	}

	showStories()
}

func registerTemplates() error {
	templateMap = make(map[string]*template.Template)

	contentMap := map[string]string{
		"story": storyTemplateContent,
		"url":   urlTemplateContent,
	}

	for name, content := range contentMap {
		tmpl, err := template.New(name).Parse(content)
		if err != nil {
			return err
		}

		templateMap[name] = tmpl
	}

	return nil
}

func showStories() {
	ids, err := news.TopStories()
	if err != nil {
		log.Println(err)
		return
	}

	for i, id := range ids {
		story, err := news.GetStory(id)
		if err != nil {
			log.Println(err)
			return
		}

		tmpl := templateMap["story"]
		if urlsOnly {
			tmpl = templateMap["url"]
		}

		if err = tmpl.Execute(os.Stdout, story.ToDisplayStory(i+1)); err != nil {
			log.Println(err)
			return
		}

		if i+1 >= numberOfArticles {
			break
		}
	}
}
