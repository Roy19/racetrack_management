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
	switch cer {
	case INVALID_ENTRY_TIME:
		return "INVALID_ENTRY_TIME"
	case INVALID_EXIT_TIME:
		return "INVALID_EXIT_TIME"
	case RACETRACK_FULL:
		return "RACETRACK_FULL"
	case SUCCESS:
		return "SUCCESS"
	default:
		return "INVALID_COMMAND_RESULT"
	}
}
