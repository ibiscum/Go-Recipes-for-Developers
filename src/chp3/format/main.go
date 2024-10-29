package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	printTime := func(format string) {
		fmt.Println(format, t.Format(format))
	}
	printTime("Mon Jan 2 15:04:05 2006")

	loc, _ := time.LoadLocation("America/Denver")
	const format = "Jan 2, 2006 at 3:04pm"
	str, _ := time.ParseInLocation(format, "Jul 9, 2012 at 5:02am", loc)
	fmt.Println(str)
	str, _ = time.ParseInLocation(format, "Jan 9, 2012 at 5:02am", loc)
	fmt.Println(str)

}
