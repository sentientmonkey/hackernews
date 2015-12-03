package hackernews

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Story struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Url      string    `json:"url"`
	Score    int       `json:"score"`
	By       string    `json:"by"`
	Time     Timestamp `json:"time"`
	Comments []int     `json:"kids"`
}

func GetStory(id int) (*Story, error) {
	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", baseUrl, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var story Story

	err = json.Unmarshal(body, &story)

	return &story, err
}
