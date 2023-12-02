package main

import (
	"fmt"
	"time"
)

func OldDateFormat(now time.Time) string {
	return now.Format("2007-01-02")
}

func NewDateFormat(now time.Time) string {
	return now.Format(time.DateOnly)
}

func NewTimeFormat(now time.Time) string {
	return now.Format(time.TimeOnly )
}

func main() {
	now := time.Now()
	fmt.Printf("Old date formatting: %s\n", OldDateFormat(now))
	fmt.Printf("New date formatting: %s\n", NewDateFormat(now))
	fmt.Printf("New time formatting: %s\n", NewTimeFormat(now))
}