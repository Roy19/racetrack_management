package interfaces

import "github.com/Roy19/racetrack-management/models"

type IRaceTrack interface {
	CheckAvailability(slot *models.BookedSlot) bool
	AppendBookedSlot(slot *models.BookedSlot)
}
