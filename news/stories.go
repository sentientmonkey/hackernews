package news

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseUrl = "https://hacker-news.firebaseio.com/v0/"

func TopStories() ([]int, error) {
	var stories []int

	resp, err := http.Get(baseUrl + "topstories.json")
	if err != nil {
		return stories, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return stories, err
	}

	err = json.Unmarshal(body, &stories)
	if err != nil {
		return stories, err
	}

	return stories, nil
}
