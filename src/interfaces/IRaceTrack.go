package interfaces

import "github.com/Roy19/racetrack-management/src/models"

type IRaceTrack interface {
	CheckAvailability(slot *models.BookedSlot) bool
	AppendBookedSlot(slot *models.BookedSlot)
}
