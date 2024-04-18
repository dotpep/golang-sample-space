package main

import (
	"fmt"
	"time"
)

// Custom Day type and Enums of week days
//type Day int

//const (
//	Sunday Day = iota
//	Monday
//	Tuesday
//	Wednesday
//	Thursday
//	Friday
//	Saturday
//)

func main() {
	now := time.Now()
	weekday := now.Weekday()

	var currendWeekday string
	var currendWeekdayStatus string

	var dayToSting = map[time.Weekday]string{
		time.Sunday:    "Sunday",
		time.Monday:    "Monday",
		time.Tuesday:   "Tuesday",
		time.Wednesday: "Wednesday",
		time.Thursday:  "Thursday",
		time.Friday:    "Friday",
		time.Saturday:  "Saturday",
	}

	switch weekday {
	case time.Sunday:
		currendWeekday = dayToSting[time.Sunday]
	case time.Monday:
		currendWeekday = dayToSting[time.Monday]
	case time.Tuesday:
		currendWeekday = dayToSting[time.Tuesday]
	case time.Wednesday:
		currendWeekday = dayToSting[time.Wednesday]
	case time.Thursday:
		currendWeekday = dayToSting[time.Thursday]
	case time.Friday:
		currendWeekday = dayToSting[time.Friday]
	case time.Saturday:
		currendWeekday = dayToSting[time.Saturday]
	}

	switch weekday {
	case time.Saturday, time.Sunday:
		currendWeekdayStatus = "It's the weekend"
	default:
		currendWeekdayStatus = "It's a weekday"
	}

	fmt.Printf("%v, %v \n", currendWeekday, currendWeekdayStatus)
	fmt.Printf("%v, %v \n", now.Format(time.DateOnly), now.Format(time.Kitchen))
}
