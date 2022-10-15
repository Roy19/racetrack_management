package services

import "github.com/Roy19/racetrack-management/models"

type BookingService struct {
	raceTrackManagement *models.RaceTrackManagement
}

func (b *BookingService) TryBookingSlot(slotToBook *models.BookedSlot) bool {
	racetracks := b.raceTrackManagement.GetRacetrackForVehicleType(slotToBook.Vehicle.VehicleType)
	for _, racetrack := range racetracks {
		if checkIfSlotCanBeAllocated(racetrack.BookedSlots, slotToBook) {
			racetrack.AppendBookedSlot(slotToBook)
			return true
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
