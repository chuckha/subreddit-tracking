package srt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var Key string

func init() {
	Key = os.Getenv("REDDIT_API_KEY")
	Key = "hi"
	if Key == "" {
		panic("Need to set REDDIT_API_KEY env var")
	}
}

const newSubredditsURL = "http://www.reddit.com/subreddits/new.json?after=%v"

type Response struct {
	Data struct {
		Children []sr
	}
	After string
}

type sr struct {
	Data Subreddit
}

func GetNewSubredditsURL(after string) string {
	return fmt.Sprintf(newSubredditsURL, after)
}

func GetSubreddits() ([]sr, error) {
	respc := DefaultAPIPool.AddURL(GetNewSubredditsURL(""))
	resp := <-respc

	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	subreddits := &Response{}
	err = json.Unmarshal(bodybytes, subreddits)
	if err != nil {
		return nil, err
	}
	return subreddits.Data.Children, nil
}
