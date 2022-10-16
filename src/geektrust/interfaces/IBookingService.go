package interfaces

import (
	"time"

	"geektrust/models"
)

type IBookingService interface {
	TryBookingSlot(slot *models.BookedSlot) bool
	AdditionalTimeForVehicle(vehicleNumber string, exitTime time.Time) bool
}
