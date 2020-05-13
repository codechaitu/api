package errors

type ExitCode int

const (
	Nil         ExitCode = 0
	Error       ExitCode = 1
	LoggerError ExitCode = 2
)
