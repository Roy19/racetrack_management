package commands

type CommandType int

const (
	BOOK CommandType = iota
	ADDITIONAL
	REVENUE
)

type CommandExecutionResult int

const (
	INVALID_ENTRY_TIME CommandExecutionResult = iota
	INVALID_EXIT_TIME
	RACETRACK_FULL
	SUCCESS
)
