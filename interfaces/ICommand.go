package interfaces

type ICommand interface {
	VerifyCommand() bool
}

type ICommandMatcher interface {
	CheckIfSame(command ICommand) bool
}
