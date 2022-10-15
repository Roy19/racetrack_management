package commands

import (
	"testing"

	"github.com/Roy19/racetrack-management/interfaces"
)

const error_format_string = "Failed in method %v. Input: %v, Got: %v, Expected: %v\n"

type baseCommandTestTable struct {
	command  interfaces.ICommand
	expected bool
}

var inputTable = []baseCommandTestTable{
	{
		Command{
			Type: "BOOK",
		},
		true,
	},
	{
		Command{
			Type: "ADDITIONAL",
		},
		true,
	},
	{
		Command{
			Type: "REVENUE",
		},
		true,
	},
	{
		Command{
			Type: "HELLO",
		},
		false,
	},
	{
		Command{
			Type: "ANOTHERCOMMAND",
		},
		false,
	},
	{
		Command{
			Type: "NEW",
		},
		false,
	},
	{
		BookCommand{
			CommandType: Command{
				"BOOK",
			},
			VehicleType:   "CAR",
			VehicleNumber: "XY4",
			EntryTime:     "13:59",
		},
		true,
	},
	{
		BookCommand{
			CommandType: Command{
				"BOOK",
			},
			VehicleType:   "BIKE",
			VehicleNumber: "XY4",
			EntryTime:     "14:59",
		},
		true,
	},
	{
		BookCommand{
			CommandType: Command{
				"BOOK",
			},
			VehicleType:   "CAR",
			VehicleNumber: "XY4",
			EntryTime:     "13:00",
		},
		true,
	},
	{
		BookCommand{
			CommandType: Command{
				"COMMAND",
			},
			VehicleType:   "CAR",
			VehicleNumber: "XY4",
			EntryTime:     "13:59",
		},
		false,
	},
	{
		BookCommand{
			CommandType: Command{
				"ADDITIONAL",
			},
			VehicleType:   "BIKE",
			VehicleNumber: "XY4",
			EntryTime:     "07:59",
		},
		false,
	},
	{
		BookCommand{
			CommandType: Command{
				"BOOK",
			},
			VehicleType:   "CAR",
			VehicleNumber: "XY4",
			EntryTime:     "12:50",
		},
		false,
	},
	{
		AdditionalCommand{
			CommandType: Command{
				"ADDITIONAL",
			},
			VehicleNumber: "XY4",
			ExitTime:      "18:40",
		},
		true,
	},
	{
		AdditionalCommand{
			CommandType: Command{
				"ADDITIONAL",
			},
			VehicleNumber: "XY4",
			ExitTime:      "15:40",
		},
		true,
	},
	{
		AdditionalCommand{
			CommandType: Command{
				"ADDITIONAL",
			},
			VehicleNumber: "XY4",
			ExitTime:      "18:70",
		},
		false,
	},
	{
		AdditionalCommand{
			CommandType: Command{
				"ADDITIONAL",
			},
			VehicleNumber: "XY4",
			ExitTime:      "20:40",
		},
		false,
	},
	{
		AdditionalCommand{
			CommandType: Command{
				"ADDITIONALCOMMAND",
			},
			VehicleNumber: "XY4",
			ExitTime:      "18:40",
		},
		false,
	},
}

func TestCommandVerifyCommand(t *testing.T) {
	for _, test := range inputTable {
		got := test.command.VerifyCommand()
		if got != test.expected {
			t.Errorf(error_format_string, "TestBaseCommandVerifyCommand", test.command, got, test.expected)
		}
	}
}
