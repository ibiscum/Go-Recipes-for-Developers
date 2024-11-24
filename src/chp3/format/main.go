package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2024, 3, 8, 18, 2, 13, 500, time.UTC)

	fmt.Println("Date in yyyy/mm/dd format", t.Format("2006/01/02"))
	fmt.Println("Date in yyyy/m/d format", t.Format("2006/1/2"))
	fmt.Println("Date in yy/m/d format", t.Format("06/1/2"))

	fmt.Println("Time in hh:mm format (12 hr)", t.Format("03:04"))
	fmt.Println("Time in hh:m format (24 hr)", t.Format("15:4"))

	fmt.Println("Date-time with time zone", t.Format("2006-01-02 13:04:05 -07:00"))

	printTime := func(format string) {
		fmt.Println(format, t.Format(format))
	}
	printTime("Mon Jan 2 15:04:05 2006")

	t, err := time.Parse("2006-01-02", "2024-08-12")
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	loc, _ := time.LoadLocation("America/Denver")
	const format = "Jan 2, 2006 at 3:04pm"
	t, err = time.ParseInLocation(format, "Jul 9, 2012 at 5:02am", loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	t, err = time.ParseInLocation(format, "Jan 9, 2012 at 5:02am", loc)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

}
