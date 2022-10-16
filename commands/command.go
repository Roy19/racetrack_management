package commands

import "time"

type Command struct {
	Type string
}

func (b Command) VerifyCommand() bool {
	return b.Type == BOOKCOMMAND || b.Type == ADDITIONALCOMMAND || b.Type == REVENUECOMMAND
}

func (b Command) CheckIfSame(command Command) bool {
	if b.Type == command.Type {
		return true
	} else {
		return false
	}
}

type BookCommand struct {
	CommandType   Command
	VehicleType   string
	VehicleNumber string
	EntryTime     string
}

func (b BookCommand) VerifyCommand() bool {
	if !b.CommandType.VerifyCommand() && b.CommandType.Type != BOOKCOMMAND {
		return false
	}
	if b.VehicleType != BIKE && b.VehicleType != CAR && b.VehicleType != SUV {
		return false
	}
	b.EntryTime += ":00"
	t, err := time.Parse("15:04:05", b.EntryTime)
	return err == nil && t.After(GetValidBookingStartTime()) && t.Before(GetValidBookingEndTime())
}

func (b BookCommand) CheckIfSame(command BookCommand) bool {
	if !b.CommandType.CheckIfSame(command.CommandType) {
		return false
	}
	if b.VehicleType != command.VehicleType {
		return false
	}
	if b.VehicleNumber != command.VehicleNumber {
		return false
	}
	if b.EntryTime != command.EntryTime {
		return false
	}
	return true
}

type AdditionalCommand struct {
	CommandType   Command
	VehicleNumber string
	ExitTime      string
}

func (a AdditionalCommand) VerifyCommand() bool {
	if !a.CommandType.VerifyCommand() && a.CommandType.Type != ADDITIONALCOMMAND {
		return false
	}
	a.ExitTime += ":00"
	t, err := time.Parse("15:04:05", a.ExitTime)
	return err == nil && t.Before(GetValidEndTime())
}

func (a AdditionalCommand) CheckIfSame(command AdditionalCommand) bool {
	if !a.CommandType.CheckIfSame(command.CommandType) {
		return false
	}
	if a.VehicleNumber != command.VehicleNumber {
		return false
	}
	if a.ExitTime != command.ExitTime {
		return false
	}
	return true
}
