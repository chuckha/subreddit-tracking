package srt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
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
		After    string
	}
}

type sr struct {
	Data Subreddit
}

func GetNewSubredditsURL(after string) string {
	return fmt.Sprintf(newSubredditsURL, after)
}

func StartGetSubreddits() {
	ticker := time.Tick(time.Duration(2) * time.Minute)
	for {
		select {
		case <-ticker:
			fmt.Println("Starting to get all subreddits")
			GetSubreddits("")
			fmt.Println("Finished getting all subreddits")
		}
	}
}

func GetSubreddits(after string) {
	fmt.Println("After: " + after)
	respc := DefaultAPIPool.AddURL(GetNewSubredditsURL(after))
	resp := <-respc
	if resp == nil {
		fmt.Println("unable to get response from reddit")
		return
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp.Body.Close()
	subreddits := &Response{}
	err = json.Unmarshal(bodybytes, subreddits)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sr := range subreddits.Data.Children {
		fmt.Println(sr.Data.Name, sr.Data.Age())
		if sr.Data.Age() > 10 {
			fmt.Println("Prune Tracking")
			fmt.Println(Tracking)
			return
		}
		Tracking[sr.Data.Name] = sr.Data
	}
	GetSubreddits(subreddits.Data.After)
}
