package main

import (
	"fmt"
	"time"
)

func daysUntilNewYear(now time.Time) int {
	newYear := time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
	return int(newYear.Sub(now).Hours() / 24)
}

func main() {
	fmt.Println(daysUntilNewYear(time.Now()))
}
