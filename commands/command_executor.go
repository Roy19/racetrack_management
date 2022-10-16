package commands

import (
	"errors"
	"fmt"
	"time"

	"github.com/Roy19/racetrack-management/interfaces"
	"github.com/Roy19/racetrack-management/models"
)

type CommandExecutor struct {
	BookingController interfaces.IBookingController
	RevenueController interfaces.IRevenueController
}

func (ce *CommandExecutor) ExecutorCommand(tokens []string) {
	_, err := createCommandAndVerify(tokens)
	if err != nil {
		fmt.Println("Invalid command")
	}
	commandType := tokens[0]
	if commandType == BOOKCOMMAND {
		vehicleType := getVehicleTypeFromString(tokens[1])
		vehicleNumber := tokens[2]
		st := tokens[3] + ":00"
		startTime, err := time.Parse("15:04:05", st)
		if err != nil {
			fmt.Printf("Failed to parse time\n")
		}
		ce.BookingController.BookSlot(vehicleType, vehicleNumber, startTime)
	} else if commandType == ADDITIONALCOMMAND {

	} else if commandType == REVENUECOMMAND {
		regular, vip := ce.RevenueController.CalculateRevenue()
		fmt.Printf("%v %v\n", regular, vip)
	}
}

func getVehicleTypeFromString(strType string) models.VehicleType {
	if strType == BIKE {
		return models.BIKE
	} else if strType == CAR {
		return models.CAR
	} else {
		return models.SUV
	}
}

func createCommandAndVerify(tokens []string) (interfaces.ICommand, error) {
	if len(tokens) == 0 {
		return nil, errors.New("empty commands not allowed")
	}
	var command interfaces.ICommand
	if tokens[0] == BOOKCOMMAND {
		command = BookCommand{
			CommandType: Command{
				Type: tokens[0],
			},
			VehicleType:   tokens[1],
			VehicleNumber: tokens[2],
			EntryTime:     tokens[3],
		}
		if !command.VerifyCommand() {
			return nil, errors.New("invalid Booking Command")
		}
	}
	if tokens[0] == ADDITIONALCOMMAND {
		command = AdditionalCommand{
			CommandType: Command{
				Type: tokens[0],
			},
			VehicleNumber: tokens[2],
			ExitTime:      tokens[3],
		}
		if !command.VerifyCommand() {
			return nil, errors.New("invalid Additional Command")
		}
	}
	if tokens[0] == REVENUECOMMAND {
		command = Command{
			Type: REVENUECOMMAND,
		}
		if !command.VerifyCommand() {
			return nil, errors.New("invalid Revenue Command")
		}
	}
	return command, nil
}
