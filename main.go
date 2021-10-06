package main

// Simply returns 0 if the sun is up, and 1 if it's down

import (
	"flag"
	"os"
	"time"

	"github.com/nathan-osman/go-sunrise"
)

func main() {
	// Default is downtown Vancouver
	lat := flag.Float64("lat", 49.28307, "Latitude")
	long := flag.Float64("long", -123.12015, "Longitude")
	flag.Parse()

	now := time.Now()
	srise, sset := sunrise.SunriseSunset(*lat, *long, now.Year(), now.Month(), now.Day())

	if now.After(sset) || now.Before(srise) {
		os.Exit(1)
	}
	os.Exit(0)

}
