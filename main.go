package main

import (
	"flag"
	"hackernews/news"
	"html/template"
	"log"
	"os"
)

const storyTemplateContent = `{{.Number | printf "%2d"}}. {{.Title}} ({{.Url}})
    {{.Score}} points by {{.By}} {{.FormattedTime}} | {{.CommentCount}} comments
`
const urlTemplateContent = `{{.Url}}
`

type App struct {
	numberOfArticles int
	urlsOnly         bool
	templateMap      map[string]*template.Template
}

func NewApp() *App {
	return &App{
		templateMap: make(map[string]*template.Template),
	}
}

func (app *App) Run() {
	app.parseFlags()
	app.registerTemplates()
	app.showStories()
}

func (app *App) parseFlags() {
	flag.IntVar(&app.numberOfArticles, "n", 10, "number of articles")
	flag.BoolVar(&app.urlsOnly, "u", false, "only show urls")

	flag.Parse()
}

func (app *App) registerTemplates() error {
	contentMap := map[string]string{
		"story": storyTemplateContent,
		"url":   urlTemplateContent,
	}

	for name, content := range contentMap {
		tmpl, err := template.New(name).Parse(content)
		if err != nil {
			log.Fatal(err)
		}

		app.templateMap[name] = tmpl
	}

	return nil
}

func (app *App) showStories() {
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

		tmpl := app.templateMap["story"]
		if app.urlsOnly {
			tmpl = app.templateMap["url"]
		}

		if err = tmpl.Execute(os.Stdout, story.ToDisplayStory(i+1)); err != nil {
			log.Println(err)
			return
		}

		if i+1 >= app.numberOfArticles {
			break
		}
	}
}

func main() {
	NewApp().Run()
}
