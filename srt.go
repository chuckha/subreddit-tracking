package srt

import (
	"math"
	"time"
)

const (
	Gravity = 1.8
)

var DefaultAPIPool = NewAPIPool()

type Subreddit struct {
	Name        string  `json:"display_name"`
	Title       string  `json:"title"`
	Subscribers int     `json:"subscribers"`
	Created     float64 `json:"created_utc"`
	NSFW        bool    `json:"over18"`
}

// Return number of hours
func (s *Subreddit) Age() float64 {
	created := time.Unix(int64(s.Created), 0)
	return time.Since(created).Hours()
}
func (s *Subreddit) Score() float64 {
	if s.Subscribers-1 <= 0 {
		return 0.0
	}
	return float64(s.Subscribers-1) / math.Pow(s.Age()+2.0, Gravity)
}
func (s *Subreddit) Update() {
	response := DefaultAPIPool.AddURL(URLForSubscriberCount(s.Name))
	<-response
}
