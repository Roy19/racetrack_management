package controllers

import (
	"time"

	"github.com/Roy19/racetrack-management/interfaces"
	"github.com/Roy19/racetrack-management/models"
)

type BookingController struct {
	bookingService interfaces.IBookingService
}

func (b *BookingController) BookSlot(vehicleType models.VehicleType,
	identificationNumber string,
	startTime time.Time) bool {
	endTime := startTime.Add(time.Hour * time.Duration(3))
	slotToBook := &models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          vehicleType,
			IdentificationNumber: identificationNumber,
		},
		StartTime: startTime,
		EndTime:   endTime,
	}
	// try booking a slot
	ans := b.bookingService.TryBookingSlot(slotToBook)
	// check if vehicle can be assigned
	return ans
}
