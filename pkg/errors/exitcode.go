package errors

type ExitCoder interface {
	error
	ExitCode() int
}

type ExitCodeError struct {
	innerErr   error
	statusCode int
}

var _ ExitCoder = ExitCodeError{}

func NewExitCode(innerErr error, statusCode int) ExitCodeError {
	return ExitCodeError{
		innerErr:   innerErr,
		statusCode: statusCode,
	}
}

func (e ExitCodeError) Error() string {
	return e.innerErr.Error()
}

func (e ExitCodeError) String() string {
	return e.Error()
}

func (e ExitCodeError) Unwrap() error {
	return e.innerErr
}

func (e ExitCodeError) ExitCode() int {
	return e.statusCode
}
