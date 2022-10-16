package models

import "time"

type Vehicle struct {
	VehicleType          VehicleType
	IdentificationNumber string
}

type BookedSlot struct {
	Vehicle   *Vehicle
	StartTime time.Time
	EndTime   time.Time
}

type RaceTrack struct {
	RaceTrackType      RacetrackType
	AllowedVehicleType VehicleType
	BookedSlots        []*BookedSlot
}

func (raceTrack *RaceTrack) CheckAvailability(slot *BookedSlot) bool {
	return slot.Vehicle.VehicleType != raceTrack.AllowedVehicleType
}

func (raceTrack *RaceTrack) AppendBookedSlot(slot *BookedSlot) {
	raceTrack.BookedSlots = append(raceTrack.BookedSlots, slot)
}

type RaceTrackManagement struct {
	RaceTracks []*RaceTrack
}

func (rcm *RaceTrackManagement) GetTotalSlotsBooked() int {
	totalSlotsBooked := 0
	for _, v := range rcm.RaceTracks {
		totalSlotsBooked += len(v.BookedSlots)
	}
	return totalSlotsBooked
}

func (rcm *RaceTrackManagement) GetRacetrackForVehicleType(vehicleType VehicleType) []*RaceTrack {
	var raceTracks []*RaceTrack
	if vehicleType == BIKE {
		for _, v := range rcm.RaceTracks {
			if v.RaceTrackType == REGULAR && v.AllowedVehicleType == BIKE {
				raceTracks = append(raceTracks, v)
			}
		}
	} else if vehicleType == CAR {
		for _, v := range rcm.RaceTracks {
			if v.RaceTrackType == REGULAR && v.AllowedVehicleType == CAR {
				raceTracks = append(raceTracks, v)
			}
			if v.RaceTrackType == VIP && v.AllowedVehicleType == CAR {
				raceTracks = append(raceTracks, v)
			}
		}
	} else {
		for _, v := range rcm.RaceTracks {
			if v.RaceTrackType == REGULAR && v.AllowedVehicleType == SUV {
				raceTracks = append(raceTracks, v)
			}
			if v.RaceTrackType == VIP && v.AllowedVehicleType == SUV {
				raceTracks = append(raceTracks, v)
			}
		}
	}

	return raceTracks
}
