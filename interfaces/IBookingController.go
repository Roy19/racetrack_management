package interfaces

import (
	"time"

	"github.com/Roy19/racetrack-management/models"
)

type IBookingController interface {
	BookSlot(vehicleType models.VehicleType,
		identificationNumber string,
		startTime time.Time) bool
}
