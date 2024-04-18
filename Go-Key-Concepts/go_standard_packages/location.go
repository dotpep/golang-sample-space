package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Current: ", time.Now())

	fmt.Println("---")

	locations := map[string]*time.Location{
		"Japan":  getLocation("Asia/Tokyo"),
		"UK":     getLocation("Europe/London"),
		"Russia": getLocation("Europe/Moscow"),
	}

	time.Local = locations["Japan"]
	fmt.Println("Japan: ", time.Now())

	time.Local = locations["UK"]
	fmt.Println("UK: ", time.Now())
}

func getLocation(location string) *time.Location {
	loc, err := time.LoadLocation(location)
	if err != nil {
		panic(fmt.Errorf("error loading location %s: %v", location, err))
	}
	return loc
}
