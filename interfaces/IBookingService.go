package interfaces

import (
	"time"

	"github.com/Roy19/racetrack-management/models"
)

type IBookingService interface {
	TryBookingSlot(slot *models.BookedSlot) bool
	AdditionalTimeForVehicle(vehicleNumber string, exitTime time.Time) bool
}
