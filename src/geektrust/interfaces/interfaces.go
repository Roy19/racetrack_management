package interfaces

import (
	"geektrust/models"
	"time"
)

// controllers
type IBookingController interface {
	BookSlot(command ICommand) CommandExecutionResult
	AdditionalTime(command ICommand) CommandExecutionResult
}

type IRevenueController interface {
	CalculateRevenue() (int, int)
}

type IRaceTrack interface {
	CheckAvailability(slot *models.BookedSlot) bool
	AppendBookedSlot(slot *models.BookedSlot)
}

type ICommand interface {
	VerifyCommand() bool
	CheckIfSame(command ICommand) bool
}

// services
type IRevenueService interface {
	CalculateRevenue() (int, int)
}

type IBookingService interface {
	TryBookingSlot(slot *models.BookedSlot) bool
	AdditionalTimeForVehicle(vehicleNumber string, exitTime time.Time) bool
}

// builders
type IRaceTrackManagementBuilder interface {
	AddRacetrackForVechicleAndRacetrackType(
		vechicleType models.VehicleType,
		racetrackType models.RacetrackType,
		times int) IRaceTrackManagementBuilder
	BuildRacetrack() models.RaceTrackManagement
}
