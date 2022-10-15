package commands

import (
	"log"
	"time"
)

func GetValidBookingStartTime() time.Time {
	t, err := time.Parse("07:00:00", "13:00:00")
	if err != nil {
		log.Fatal("Failed to get valid start time")
	}
	return t
}

func GetValidBookingEndTime() time.Time {
	t, err := time.Parse("07:00:00", "17:00:00")
	if err != nil {
		log.Fatal("Failed to get valid start time")
	}
	return t
}

func GetValidEndTime() time.Time {
	t, err := time.Parse("07:00:00", "20:00:00")
	if err != nil {
		log.Fatal("Failed to get valid start time")
	}
	return t
}
