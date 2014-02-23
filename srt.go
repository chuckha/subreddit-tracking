package srt

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

const (
	Gravity  = 1.8
	AboutURL = "http://www.reddit.com/r/%v/about.json"
)

func URLForSubscriberCount(name string) string {
	return fmt.Sprintf(AboutURL, name)
}

var DefaultAPIPool = NewAPIPool()

type Subreddit struct {
	Name        string
	Subscribers int
	Created     time.Time
}

// Return number of hours
func (s *Subreddit) Age() float64 {
	return time.Since(s.Created).Hours()
}
func (s *Subreddit) Score() float64 {
	return float64(s.Subscribers-1) / math.Pow(s.Age()+2.0, Gravity)
}
func (s *Subreddit) Update() {
	response := make(chan *http.Response)
	DefaultAPIPool.AddCall(URLForSubscriberCount(s.Name), response)
	<-response
}

type APIPool struct {
	Calls int
	queue chan string
}

func NewAPIPool() *APIPool {
	return &APIPool{}
}

func (a *APIPool) AddCall(URL string, resp chan *http.Response) {}
func (a *APIPool) FillPool()                                    {}
func (a *APIPool) DrainPool()                                   {}
