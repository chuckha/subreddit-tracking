package srt

import (
	"encoding/json"
	"errors"
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

// Frequency in minutes
func StartGetSubreddits(frequency int) {
	ticker := time.Tick(time.Duration(frequency) * time.Minute)
	for {
		select {
		case <-ticker:
			fmt.Println("Starting to get all subreddits")
			var err error
			var after string
			for err == nil {
				after, err = getSubreddits(after)
			}
			fmt.Println(err)
			fmt.Println("Finished getting all subreddits")
		}
	}
}

// Gets a set of subreddits and returns the "after" (the code to get the next set of subreddits).
// Returns an error when there are no more or errors.
func getSubreddits(after string) (string, error) {
	respc := DefaultAPIPool.AddURL(GetNewSubredditsURL(after))
	resp := <-respc
	if resp == nil {
		return "", errors.New("unable to get a response from Reddit")
	}
	bodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	subreddits := &Response{}
	err = json.Unmarshal(bodybytes, subreddits)
	if err != nil {
		return "", err
	}
	for _, sr := range subreddits.Data.Children {
		fmt.Println(sr.Data.Name, sr.Data.Age())
		if sr.Data.Age() > 10 {
			fmt.Println("Prune Tracking")
			return "", errors.New("end of subreddits")
		}
		Tracking[sr.Data.Name] = sr.Data
	}
	return subreddits.Data.After, nil
}
