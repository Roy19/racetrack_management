package controllers

import (
	"time"

	"github.com/Roy19/racetrack-management/commands"
	"github.com/Roy19/racetrack-management/interfaces"
	"github.com/Roy19/racetrack-management/models"
)

type BookingController struct {
	bookingService interfaces.IBookingService
}

func (b *BookingController) BookSlot(command commands.BookCommand) interfaces.CommandExecutionResult {
	if !command.VerifyCommand() {
		return interfaces.INVALID_ENTRY_TIME
	}
	command.EntryTime += ":00"
	startTime, _ := time.Parse("15:04:05", command.EntryTime)
	endTime := startTime.Add(time.Hour * time.Duration(3))
	slotToBook := &models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          getVehicleTypeFromString(command.VehicleType),
			IdentificationNumber: command.VehicleNumber,
		},
		StartTime: startTime,
		EndTime:   endTime,
	}
	if !b.bookingService.TryBookingSlot(slotToBook) {
		return interfaces.RACETRACK_FULL
	} else {
		return interfaces.SUCCESS
	}
}

func (b *BookingController) AdditionalTime(command commands.AdditionalCommand) interfaces.CommandExecutionResult {
	if !command.VerifyCommand() {
		return interfaces.INVALID_EXIT_TIME
	}
	command.ExitTime += ":00"
	exitTime, _ := time.Parse("15:04:05", command.ExitTime)
	if !b.bookingService.AdditionalTimeForVehicle(command.VehicleNumber, exitTime) {
		return interfaces.FAILURE
	} else {
		return interfaces.SUCCESS
	}
}

func getVehicleTypeFromString(strType string) models.VehicleType {
	if strType == commands.BIKE {
		return models.BIKE
	} else if strType == commands.CAR {
		return models.CAR
	} else {
		return models.SUV
	}
}
