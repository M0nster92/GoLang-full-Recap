package main

import (
	"fmt"
	"time"
)

var countryTz = map[string]string{
	"Hungary": "Europe/Budapest",
	"Egypt":   "Africa/Cairo",
	"Canada":  "America/Toronto",
}

func timeIn(name string) time.Time {
	loc, err := time.LoadLocation(countryTz[name])
	if err != nil {
		panic(err)
	}
	return time.Now().In(loc)
}

func main() {
	utc := time.Now().UTC()
	hun := timeIn("Hungary")
	eg := timeIn("Egypt")
	canada := timeIn("Canada")
	fmt.Println("UTC ", utc)
	fmt.Println("Hungary", hun)
	fmt.Println("Egypt", eg)
	fmt.Println("Canada ", canada)
}
