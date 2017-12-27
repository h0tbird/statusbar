package main

//-----------------------------------------------------------------------------
// Imports:
//-----------------------------------------------------------------------------

import "fmt"
import "time"
import "bytes"

//-----------------------------------------------------------------------------
// Item structure:
//-----------------------------------------------------------------------------

type item struct {
	show bool
	data bytes.Buffer
	fn   func(*item)
}

func (i *item) runFunc() {
	if i.fn != nil {
		i.fn(i)
	}
}

//-----------------------------------------------------------------------------
// Item functions:
//-----------------------------------------------------------------------------

func printOne(i *item) {
	for _ = range time.NewTicker(1 * time.Second).C {
		i.data.WriteString("1")
	}
}

func printTwo(i *item) {
	for _ = range time.NewTicker(2 * time.Second).C {
		i.data.WriteString("2")
	}
}

func printFour(i *item) {
	for _ = range time.NewTicker(4 * time.Second).C {
		i.data.WriteString("4")
	}
}

//-----------------------------------------------------------------------------
// Main loop:
//-----------------------------------------------------------------------------

func main() {

	// Initialize the structure:
	items := map[string]*item{
		"one":  &item{true, *bytes.NewBufferString("1"), printOne},
		"two":  &item{true, *bytes.NewBufferString("2"), printTwo},
		"four": &item{true, *bytes.NewBufferString("4"), printFour},
	}

	// Start each item logic:
	for _, v := range items {
		go func(i *item) {
			i.runFunc()
		}(v)
	}

	// Refresh the status bar:
	for _ = range time.NewTicker(time.Second).C {
		fmt.Printf("\033[H\033[2J")
		for _, item := range items {
			if item.show {
				fmt.Println(item.data.String())
			}
		}
	}
}
