package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	newYear := time.Date(now.Year()+1, time.January, 1, 0, 0, 0, 0, now.Location())
	days := int(newYear.Sub(now).Hours() / 24)
	fmt.Println(days)
}
