package interfaces

type CommandExecutionResult int

const (
	INVALID_ENTRY_TIME CommandExecutionResult = iota
	INVALID_EXIT_TIME
	RACETRACK_FULL
	SUCCESS
	FAILURE
)

func (cer CommandExecutionResult) GetStringResult() string {
	if cer == INVALID_ENTRY_TIME {
		return "INVALID_ENTRY_TIME"
	}
	if cer == INVALID_EXIT_TIME {
		return "INVALID_EXIT_TIME"
	}
	if cer == RACETRACK_FULL {
		return "RACETRACK_FULL"
	}
	if cer == SUCCESS {
		return "SUCCESS"
	}
	return "INVALID_COMMAND_RESULT"
}
