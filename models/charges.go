package models

const (
	BIKETRACKREGULARCOSTPERHOUR int = 60
	CARTRACKREGULARCOSTPERHOUR  int = 120
	CARTRACKVIPCOSTPERHOUR      int = 250
	SUVTRACKREGULARCOSTPERHOUR  int = 200
	SUVTRACKVIPCOSTPERHOUR      int = 300
)

func GetChargeGivenTrackAndVehicleType(vehicleType VehicleType, trackType RacetrackType) int {
	if trackType == REGULAR {
		if vehicleType == BIKE {
			return BIKETRACKREGULARCOSTPERHOUR
		}
		if vehicleType == CAR {
			return CARTRACKREGULARCOSTPERHOUR
		}
		if vehicleType == SUV {
			return SUVTRACKREGULARCOSTPERHOUR
		}
	} else if trackType == VIP {
		if vehicleType == CAR {
			return CARTRACKVIPCOSTPERHOUR
		}
		if vehicleType == SUV {
			return SUVTRACKVIPCOSTPERHOUR
		}
	}
	return 0
}
