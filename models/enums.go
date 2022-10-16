package models

type RacetrackType int

const (
	REGULAR RacetrackType = iota
	VIP
)

type VehicleType int

const (
	BIKE VehicleType = iota
	CAR
	SUV
)
