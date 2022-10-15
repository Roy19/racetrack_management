package commands

import "time"

type Command struct {
	Type string
}

func (c Command) VerifyCommand() bool {
	return c.Type == BOOKCOMMAND || c.Type == ADDITIONALCOMMAND || c.Type == REVENUECOMMAND
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
