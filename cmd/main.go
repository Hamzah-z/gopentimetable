package main

import (
	"fmt"
	"time"

	"github.com/snabb/isoweek"
)

var Weekdays = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

var ReqHeaders = map[string]string{
	"Authorization": "basic T64Mdy7m[",
	"Content-Type":  "application/json; charset=utf-8",
	"credentials":   "include",
	"Referer":       "https://opentimetable.dcu.ie/",
	"Origin":        "https://opentimetable.dcu.ie/",
}

func main() {
	fmt.Println(startOfWeek("2021-03-31T15:04:05Z"))
}

func startOfWeek(Date string) string {
	dateToFetch, _ := time.Parse(time.RFC3339, Date)
	_, weekNum := dateToFetch.ISOWeek()
	firstDayInWeek := isoweek.StartTime(dateToFetch.Year(), weekNum, time.Local)

	return firstDayInWeek.Format(time.RFC3339)
}
