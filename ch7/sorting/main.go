package main

import (
	"fmt"
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }

var tracks = []*Track{
	{"Go", "Jason", "LIFO", 2022, length("3m00s")},
	{"Go", "Aason", "LIFO", 2032, length("3m00s")},
	{"Go", "Bason", "LIFO", 2092, length("3m00s")},
	{"Go", "Zason", "LIFO", 2012, length("3m00s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }

func main() {
	fmt.Printf("%v\n", tracks)
	sort.Sort(customSort{tracks, func(x, y *Track) bool { return x.Year > y.Year }})
	fmt.Printf("%v\n", tracks)
}
