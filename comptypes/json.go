package comptypes

import (
	"encoding/json"
	"fmt"
	"log"
)

func JsonBasics() {
	fmt.Println("Hello from JSON")

	data, err := json.MarshalIndent(movies, "", "  ")

	if err != nil {
		log.Fatalf("Json marshaling failed: %s", err)
	}
	fmt.Printf("%s", data)

	// unmarshaling:
	var titles []struct{ Title string }

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("Json unmarshaling failed: %s", err)
	}

	fmt.Println(titles)
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}
