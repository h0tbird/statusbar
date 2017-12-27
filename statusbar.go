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
	iconBattery   = "\xEE\x80\xB3"
	iconUpdates   = "\xEE\x80\x8E"
)

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
		setStatus(strings.Join(s.data, fieldSeparator) + "          ")
	}
}

//-----------------------------------------------------------------------------
// Item functions:
//-----------------------------------------------------------------------------

func updates(i *item) {
	for {
		i.data = softWhite + iconUpdates + "0"
		time.Sleep(5 * time.Minute)
	}
}

func battery(i *item) {
	for {
		i.data = softPurple + iconBattery + "100%"
		time.Sleep(time.Minute)
	}
}

func dateTime(i *item) {
	for {
		i.data = time.Now().Format(softWhite + iconWallClock +
			"Mon _2 Jan " + iconLeftArrow + "15:04:05")
		time.Sleep(time.Second)
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
