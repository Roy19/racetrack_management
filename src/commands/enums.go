package commands

type CommandType string

const (
	BOOK       CommandType = "BOOK"
	ADDITIONAL CommandType = "ADDITIONAL"
	REVENUE    CommandType = "REVENUE"
)

type CommandExecutionResult string

const (
	INVALID_ENTRY_TIME CommandExecutionResult = "INVALID_ENTRY_TIME"
	INVALID_EXIT_TIME  CommandExecutionResult = "INVALID_EXIT_TIME"
	RACETRACK_FULL     CommandExecutionResult = "RACETRACK_FULL"
	SUCCESS            CommandExecutionResult = "SUCCESS"
)
