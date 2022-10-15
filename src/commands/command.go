package commands

import "time"

type Command struct {
	Type string
}

func (c Command) VerifyCommand() bool {
	return c.Type == "BOOK" || c.Type == "ADDITIONAL" || c.Type == "REVENUE"
}

type BookCommand struct {
	CommandType   Command
	VehicleType   string
	VehicleNumber string
	EntryTime     string
}

func (b BookCommand) VerifyCommand() bool {
	if !b.CommandType.VerifyCommand() {
		return false
	}
	if b.VehicleType != "BIKE" && b.VehicleType != "CAR" && b.VehicleType != "SUV" {
		return false
	}
	b.EntryTime += ":00"
	t, err := time.Parse("07:00:00", b.EntryTime)
	return err == nil && t.After(GetValidBookingStartTime()) && t.Before(GetValidBookingEndTime())
}

type AdditionalCommand struct {
	CommandType   Command
	VehicleNumber string
	ExitTime      string
}

func (a AdditionalCommand) VerifyCommand() bool {
	if !a.CommandType.VerifyCommand() {
		return false
	}
	a.ExitTime += ":00"
	t, err := time.Parse("07:00:00", a.ExitTime)
	return err == nil && t.Before(GetValidEndTime())
}
