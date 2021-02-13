package errorlib

type (
	CustomError struct {
		error
		Traces    []string
		errString string
	}
)
