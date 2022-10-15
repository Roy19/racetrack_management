package interfaces

import "github.com/Roy19/racetrack-management/models"

type IBookingService interface {
	TryBookingSlot(slot *models.BookedSlot) bool
}
