package models

type RacetrackType int

const (
	REGULAR RacetrackType = iota
	VIP
)

type VehicleType string

const (
	BIKE VehicleType = "BIKE"
	CAR  VehicleType = "CAR"
	SUV  VehicleType = "SUV"
)
