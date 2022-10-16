package interfaces

type ICommand interface {
	VerifyCommand() bool
	CheckIfSame(command ICommand) bool
}
