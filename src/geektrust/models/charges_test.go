package models

import (
	"testing"
)

type chargesTestTable struct {
	vehicleType   VehicleType
	raceTrackType RacetrackType
	chargePerHour int
}

var (
	inputTable []chargesTestTable = []chargesTestTable{
		{
			vehicleType:   BIKE,
			raceTrackType: REGULAR,
			chargePerHour: BIKETRACKREGULARCOSTPERHOUR,
		},
		{
			vehicleType:   CAR,
			raceTrackType: REGULAR,
			chargePerHour: CARTRACKREGULARCOSTPERHOUR,
		},
		{
			vehicleType:   SUV,
			raceTrackType: REGULAR,
			chargePerHour: SUVTRACKREGULARCOSTPERHOUR,
		},
		{
			vehicleType:   CAR,
			raceTrackType: VIP,
			chargePerHour: CARTRACKVIPCOSTPERHOUR,
		},
		{
			vehicleType:   SUV,
			raceTrackType: VIP,
			chargePerHour: SUVTRACKVIPCOSTPERHOUR,
		},
	}
	error_format_string = "Failed in method %v. Input: %v, Got: %v, Expected: %v\n"
)

func TestGetChargeGivenTrackAndVehicleType(t *testing.T) {
	for _, v := range inputTable {
		got := GetChargeGivenTrackAndVehicleType(v.vehicleType, v.raceTrackType)
		expected := v.chargePerHour
		if got != expected {
			t.Errorf(error_format_string, "TestGetChargeGivenTrackAndVehicleType", v, got, expected)
		}
	}
}
