package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

func main() {
	movies := []Movie{
		{Title: "Title1", Year: 1990, Color: false, Actors: []string{"Jason1, Jason2"}},
		{Title: "Title2", Year: 1992, Color: true, Actors: []string{"Jason4, Jason6"}},
		{Title: "Title3", Year: 1994, Color: true, Actors: []string{"Jason8, Jason10"}},
		{Title: "Title4", Year: 1996, Color: false, Actors: []string{"Jason12, Jason14"}},
		{Title: "Title5", Color: true, Actors: []string{"Jason18, Jason20"}},
	}
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	var titles []Movie
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Printf("%v\n", titles)
}
