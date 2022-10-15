package models

import "time"

type Vehicle struct {
	VehicleType          VehicleType
	IdentificationNumber string
}

type RaceTrack struct {
	RaceTrackType      RacetrackType
	AllowedVehicleType VehicleType
	BookedSlots        []*BookedSlot
}

type BookedSlot struct {
	Vehicle  *Vehicle
	TimeSlot time.Time
}

type RaceTrackManagement struct {
	RaceTracks []*RaceTrack
}

func (raceTrack *RaceTrack) CheckAvailability(slot *BookedSlot) bool {
	return slot.Vehicle.VehicleType != raceTrack.AllowedVehicleType
}

func (raceTrack *RaceTrack) AppendBookedSlot(slot *BookedSlot) {
	raceTrack.BookedSlots = append(raceTrack.BookedSlots, slot)
}
