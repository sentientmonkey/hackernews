package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"time"

	"github.com/sentientmonkey/hackernews/news"
)

const storyTemplateContent = `{{.Number | printf "%2d"}}. {{.Title}} ({{.Url}})
    {{.Score}} points by {{.By}} {{.FormattedTime}} | {{.CommentCount}} comments
`
const urlTemplateContent = `{{.Url}}
`

type App struct {
	numberOfArticles int
	urlsOnly         bool
	verbose          bool
	templateMap      map[string]*template.Template
	fetchChannel     chan *StoryId
	printChannel     chan *news.DisplayStory
	doneChannel      chan bool
}

type StoryId struct {
	Id     int64
	Number int
}

const maxChanSize = 10

func NewApp() *App {
	return &App{
		templateMap:  make(map[string]*template.Template),
		fetchChannel: make(chan *StoryId, maxChanSize),
		printChannel: make(chan *news.DisplayStory, maxChanSize),
		doneChannel:  make(chan bool),
	}
}

func (app *App) Run() {
	startTime := time.Now()
	app.parseFlags()
	app.registerTemplates()
	app.showStories()

	if app.verbose {
		fmt.Printf("Took %s\n", time.Since(startTime).String())
	}
}

func (app *App) parseFlags() {
	flag.IntVar(&app.numberOfArticles, "n", 10, "number of articles")
	flag.BoolVar(&app.urlsOnly, "u", false, "only show urls")
	flag.BoolVar(&app.verbose, "v", false, "verbose")

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

	go app.storyFetcher()
	go app.storyPrinter()

	for i, id := range ids {
		app.fetchChannel <- &StoryId{Number: i + 1, Id: id}

		if i+1 >= app.numberOfArticles {
			break
		}
	}

	close(app.fetchChannel)

	<-app.doneChannel
}

func (app *App) storyFetcher() {
	for storyId := range app.fetchChannel {
		story, err := news.GetStory(storyId.Id)
		if err != nil {
			log.Println(err)
			break
		}

		app.printChannel <- story.ToDisplayStory(storyId.Number)
	}
	close(app.printChannel)
}

func (app *App) storyPrinter() {
	tmpl := app.templateMap["story"]
	if app.urlsOnly {
		tmpl = app.templateMap["url"]
	}

	for displayStory := range app.printChannel {
		if err := tmpl.Execute(os.Stdout, displayStory); err != nil {
			log.Println(err)
			break
		}
	}

	app.doneChannel <- true
}

func main() {
	NewApp().Run()
}
