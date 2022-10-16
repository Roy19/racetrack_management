package commands

import (
	"errors"
	"fmt"

	"github.com/Roy19/racetrack-management/interfaces"
)

type CommandExecutor struct {
	BookingController interfaces.IBookingController
	RevenueController interfaces.IRevenueController
}

func (ce *CommandExecutor) ExecutorCommand(tokens []string) {
	command, err := createCommandAndVerify(tokens)
	if err != nil {
		fmt.Println("Invalid command")
	}
	commandType := tokens[0]
	if commandType == BOOKCOMMAND {
		result := ce.BookingController.BookSlot(command)
		fmt.Println(result.GetStringResult())
	} else if commandType == ADDITIONALCOMMAND {
		result := ce.BookingController.AdditionalTime(command)
		fmt.Println(result.GetStringResult())
	} else if commandType == REVENUECOMMAND {
		regular, vip := ce.RevenueController.CalculateRevenue()
		fmt.Printf("%v %v\n", regular, vip)
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
			return nil, errors.New("invalid booking command")
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
			return nil, errors.New("invalid additional command")
		}
	}
	if tokens[0] == REVENUECOMMAND {
		command = Command{
			Type: REVENUECOMMAND,
		}
		if !command.VerifyCommand() {
			return nil, errors.New("invalid revenue command")
		}
	}
	return command, nil
}
