package srt

import (
	"testing"
)

var agetc = []struct {
	sr    *Subreddit
	age   float64 // just be older than this number
	score float64 // also monotonic decreasing so just be smaller
}{
	{
		sr: &Subreddit{
			Name:        "golang",
			Subscribers: 7274,
			Created:     1257900868.0,
		},
		age:   37575.7,
		score: 4.2348429747663435e-05,
	},
	{
		sr: &Subreddit{
			Name:        "funny",
			Subscribers: 5403592,
			Created:     1201242956.0,
		},
		age:   53314.1,
		score: 0.016762379635540822,
	},
	{
		sr: &Subreddit{
			Name:        "hihi",
			Subscribers: 0,
			Created:     1.39317414E9,
		},
		age:   .02,
		score: 0.0,
	},
}

func TestSubredditScore(t *testing.T) {
	for _, tc := range agetc {
		if tc.sr.Score() > tc.score {
			t.Errorf("Exepcted %v to have at most %v points", tc.sr, tc.score)
		}
	}
}

func TestSubredditAge(t *testing.T) {
	for _, tc := range agetc {
		if tc.sr.Age() < tc.age {
			t.Errorf("Exepcted %v to be at least %v old: %v", tc.sr, tc.age, tc.sr.Age())
		}
	}
}

var urltc = []struct {
	in, out string
}{
	{
		in:  "golang",
		out: "http://www.reddit.com/r/golang/about.json",
	},
	{
		in:  "",
		out: "http://www.reddit.com/r//about.json",
	},
}

func TestURLForSubscriberCount(t *testing.T) {
	for _, tc := range urltc {
		actual := URLForSubscriberCount(tc.in)
		if actual != tc.out {
			t.Errorf("Expected: %v\nGot: %v\n", tc.out, actual)
		}
	}
}
