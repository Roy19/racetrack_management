package interfaces

import "geektrust/models"

type IRaceTrack interface {
	CheckAvailability(slot *models.BookedSlot) bool
	AppendBookedSlot(slot *models.BookedSlot)
}
