package controllers

import (
	"time"

	"geektrust/commands"
	"geektrust/interfaces"
	"geektrust/models"
)

type BookingController struct {
	BookingService interfaces.IBookingService
}

func (b *BookingController) BookSlot(command interfaces.ICommand) interfaces.CommandExecutionResult {
	c := command.(commands.BookCommand)
	if !c.VerifyCommand() {
		return interfaces.INVALID_ENTRY_TIME
	}
	c.EntryTime += ":00"
	startTime, _ := time.Parse("15:04:05", c.EntryTime)
	endTime := startTime.Add(time.Hour * time.Duration(3))
	slotToBook := &models.BookedSlot{
		Vehicle: &models.Vehicle{
			VehicleType:          getVehicleTypeFromString(c.VehicleType),
			IdentificationNumber: c.VehicleNumber,
		},
		StartTime: startTime,
		EndTime:   endTime,
	}
	if !b.BookingService.TryBookingSlot(slotToBook) {
		return interfaces.RACETRACK_FULL
	} else {
		return interfaces.SUCCESS
	}
}

func (b *BookingController) AdditionalTime(command interfaces.ICommand) interfaces.CommandExecutionResult {
	c := command.(commands.AdditionalCommand)
	if !c.VerifyCommand() {
		return interfaces.INVALID_EXIT_TIME
	}
	c.ExitTime += ":00"
	exitTime, _ := time.Parse("15:04:05", c.ExitTime)
	if !b.BookingService.AdditionalTimeForVehicle(c.VehicleNumber, exitTime) {
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
