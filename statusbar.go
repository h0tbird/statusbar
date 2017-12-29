package main

//-----------------------------------------------------------------------------
// Imports:
//-----------------------------------------------------------------------------

import (
	"strings"
	"time"
)

//-----------------------------------------------------------------------------
// Constants:
//-----------------------------------------------------------------------------

const (

	// Colors:
	softWhite  = "\x02"
	softOrange = "\x03"
	softPurple = "\x06"

	// Items:
	fieldSeparator = softOrange + " | "

	// Icons:
	iconLeftArrow = "\xEE\x86\xAC"
	iconWallClock = "\xEE\x80\x95"
	iconUpdates   = "\xEE\x80\x8E"
)

//-----------------------------------------------------------------------------
// format:
//-----------------------------------------------------------------------------

func format(data []string) string {
	out := strings.Join(data, fieldSeparator)
	len := strings.Count(out, softWhite) +
		strings.Count(out, softOrange) +
		strings.Count(out, softPurple)
	return out + strings.Repeat(" ", 2*len)
}

//-----------------------------------------------------------------------------
// Item structure:
//-----------------------------------------------------------------------------

type item struct {
	show bool
	data string
	fn   func(*item)
}

func (i *item) runFunc() {
	if i.fn != nil {
		i.fn(i)
	}
}

//-----------------------------------------------------------------------------
// Status structure:
//-----------------------------------------------------------------------------

type status struct {
	data  []string
	items []*item
}

func (s *status) start() {
	for _, v := range s.items {
		go v.runFunc()
	}
}

func (s *status) refresh(d time.Duration) {
	for _ = range time.NewTicker(d).C {
		s.data = []string{}
		for _, item := range s.items {
			if item.show {
				s.data = append(s.data, item.data)
			}
		}
		setStatus(format(s.data))
	}
}

//-----------------------------------------------------------------------------
// Main loop:
//-----------------------------------------------------------------------------

func main() {

	s := status{
		items: []*item{
			&item{true, "", updates},
			&item{true, "", battery},
			&item{true, "", dateTime},
		},
	}

	s.start()
	s.refresh(time.Second)
}
