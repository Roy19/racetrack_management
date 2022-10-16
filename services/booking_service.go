package services

import "github.com/Roy19/racetrack-management/models"

type BookingService struct {
	raceTrackManagement *models.RaceTrackManagement
}

func (b *BookingService) TryBookingSlot(slotToBook *models.BookedSlot) bool {
	racetracks := b.raceTrackManagement.GetRacetrackForVehicleType(slotToBook.Vehicle.VehicleType)
	return tryToBookSlotType(racetracks, models.REGULAR, slotToBook) ||
		tryToBookSlotType(racetracks, models.VIP, slotToBook)
}

func tryToBookSlotType(racetracks []*models.RaceTrack, trackType models.RacetrackType,
	slotToBook *models.BookedSlot) bool {
	for _, racetrack := range racetracks {
		if racetrack.RaceTrackType == trackType &&
			racetrack.AllowedVehicleType == slotToBook.Vehicle.VehicleType {
			if checkIfSlotCanBeAllocated(racetrack.BookedSlots, slotToBook) {
				racetrack.AppendBookedSlot(slotToBook)
				return true
			}
		}
	}
	return false
}

func checkIfSlotCanBeAllocated(slots []*models.BookedSlot, slotToBook *models.BookedSlot) bool {
	if len(slots) == 0 {
		return true
	}
	// check overlapping slots here
	for _, slot := range slots {
		if slotToBook.StartTime.Before(slot.EndTime) && slotToBook.EndTime.After(slot.StartTime) {
			return false
		}
	}
	return true
}
