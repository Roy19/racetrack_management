package interfaces

type IBookingController interface {
	BookSlot(command ICommand) CommandExecutionResult
	AdditionalTime(command ICommand) CommandExecutionResult
}
