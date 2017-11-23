package cloudgoerror

import "errors"

var (
	ErrNeedImplement = errors.New("this function need to be implemented")

	// Time
	// InvalidTime         = errors.New("startTime/EndTime is not valid")
	ErrInvalidTimeInterval    = errors.New("the EndTime must be after StartTime")
	ErrConflictedTimeInterval = errors.New("given time interval conflicts with existed interval")

	// Information
	ErrGivenConflictedInfo = errors.New("given a not reasonable information")
)

type CloudgoError struct {
	msg string
}

func NewCloudgoError(msg string) *CloudgoError {
	return &CloudgoError{
		msg: msg,
	}
}

func (e *CloudgoError) Error() string {
	return e.msg
}
