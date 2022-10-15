package commands

import (
	"log"
	"time"
)

const (
	BIKE              string = "BIKE"
	CAR               string = "CAR"
	SUV               string = "SUV"
	BOOKCOMMAND       string = "BOOK"
	ADDITIONALCOMMAND string = "ADDITIONAL"
	REVENUECOMMAND    string = "REVENUE"
)

func GetValidBookingStartTime() time.Time {
	t, err := time.Parse("15:04:05", "12:59:59")
	if err != nil {
		log.Fatal("Failed to get valid start time")
	}
	return t
}

func GetValidBookingEndTime() time.Time {
	t, err := time.Parse("15:04:05", "17:00:01")
	if err != nil {
		log.Fatal("Failed to get valid start time")
	}
	return t
}

func GetValidEndTime() time.Time {
	t, err := time.Parse("15:04:05", "20:00:01")
	if err != nil {
		log.Fatal("Failed to get valid start time")
	}
	return t
}
