package main

import (
	"fmt"
	"time"
)

func main() {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)

	fmt.Println("Hi!")
	fmt.Printf("Today is %s (%s)\n", now.Format("2006/01/02"), now.Weekday())
}
