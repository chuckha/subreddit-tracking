package srt

import (
	"fmt"
	"net/http"
	"time"
)

const (
	AboutURL = "http://www.reddit.com/r/%v/about.json"
	// In milliseconds
	// TODO: Change me to 2000 since we have 1call every 2 seconds
	APIRate = 2000
)

type Call struct {
	URL  string
	resp chan *http.Response
}

func NewCall(URL string) *Call {
	return &Call{
		URL:  URL,
		resp: make(chan *http.Response),
	}
}

func URLForSubscriberCount(name string) string {
	return fmt.Sprintf(AboutURL, name)
}

type APIPool struct {
	calls chan struct{}
	queue chan *Call
}

func NewAPIPool() *APIPool {
	apipool := &APIPool{
		// Using a buffered channel will allow for bursts but will not
		// exceed the 1 request / 2 second average.
		calls: make(chan struct{}, 30),
		queue: make(chan *Call),
	}
	go apipool.FillPool()
	go apipool.DrainPool()
	return apipool
}

// Add an API call to the queue
func (a *APIPool) AddURL(URL string) chan *http.Response {
	c := NewCall(URL)
	go a.addCall(c)
	return c.resp
}

func (a *APIPool) addCall(c *Call) {
	a.queue <- c
}

// Increase the number of available api calls we have
func (a *APIPool) inc() {
	if len(a.calls) < 30 {
		a.calls <- struct{}{}
	}
}

// Add a new API call ever 500 ms
func (a *APIPool) FillPool() {
	ticker := time.Tick(time.Duration(APIRate) * time.Millisecond)
	for {
		select {
		case <-ticker:
			a.inc()
		}
	}
}

// for each item on the queue, pull out a call in the pool and make the API call
func (a *APIPool) DrainPool() {
	for {
		select {
		case call := <-a.queue:
			<-a.calls
			resp, err := http.Get(call.URL)
			if err != nil {
				close(call.resp)
				continue
			}
			call.resp <- resp
			close(call.resp)
		}
	}
}
