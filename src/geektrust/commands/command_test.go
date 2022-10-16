package commands

import (
	"errors"
	"testing"

	"geektrust/interfaces"
)

const error_format_string = "Failed in method %v. Input: %v, Got: %v, Expected: %v\n"

type baseCommandTestTable struct {
	command  interfaces.ICommand
	expected bool
}

type createCommandAndVerifyTable struct {
	tokens []string
	result interfaces.ICommand
	err    error
}

var (
	inputTableCommand = []baseCommandTestTable{
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
	inputTableCreateCommandAndVerify = []createCommandAndVerifyTable{
		{
			tokens: []string{"BOOK", "BIKE", "M40", "14:00"},
			result: BookCommand{
				CommandType: Command{
					Type: "BOOK",
				},
				VehicleType:   "BIKE",
				VehicleNumber: "M40",
				EntryTime:     "14:00",
			},
			err: nil,
		},
		{
			tokens: []string{"ADDITIONAL", "M40", "17:40"},
			result: AdditionalCommand{
				CommandType: Command{
					Type: "ADDITIONAL",
				},
				VehicleNumber: "M40",
				ExitTime:      "17:40",
			},
			err: nil,
		},
		{
			tokens: []string{"REVENUE"},
			result: Command{
				Type: "REVENUE",
			},
			err: nil,
		},
		{
			tokens: []string{"BOOK", "BIKE", "M40"},
			result: nil,
			err:    errors.New("invalid booking command"),
		},
		{
			tokens: []string{"ADDITIONAL", "20:50"},
			result: nil,
			err:    errors.New("invalid additional command"),
		},
		{
			tokens: []string{},
			result: nil,
			err:    errors.New("empty commands not allowed"),
		},
		{
			tokens: []string{"BOOK", "A66", "11:00"},
			result: nil,
			err:    errors.New("invalid booking command"),
		},
	}
)

func TestCommandVerifyCommand(t *testing.T) {
	for _, test := range inputTableCommand {
		got := test.command.VerifyCommand()
		if got != test.expected {
			t.Errorf(error_format_string, "TestBaseCommandVerifyCommand", test.command, got, test.expected)
		}
	}
}

func TestCreateCommandAndVerify(t *testing.T) {
	for _, test := range inputTableCreateCommandAndVerify {
		gotres, goterr := createCommandAndVerify(test.tokens)
		if !compareErrors(goterr, test.err) || !compareResult(gotres, test.result) {
			t.Errorf(error_format_string, "TestCreateCommandAndVerify", test.tokens, gotres, test.result)
			t.Errorf(error_format_string, "TestCreateCommandAndVerify", test.tokens, goterr, test.err)
		}
	}
}

func compareResult(got, want interfaces.ICommand) bool {
	if got == nil && want == nil {
		return true
	}
	if got != nil && want != nil && got.CheckIfSame(want) {
		return true
	}
	return false
}

func compareErrors(got, want error) bool {
	if got == nil && want == nil {
		return true
	}
	if got != nil && want != nil && got.Error() == want.Error() {
		return true
	}
	return false
}
