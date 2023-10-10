package errors

// ErrorType defines types of errors that are possible from soup
type ErrorType int

const (
	// ErrUnableToParse will be returned when the HTML could not be parsed
	ErrUnableToParse ErrorType = iota
	// ErrCreatingGetRequest will be returned when the get request couldn't be created
	ErrCreatingGetRequest
	// ErrInGetRequest will be returned when there was an error during the get request
	ErrInGetRequest
	// ErrReadingResponse will be returned if there was an error reading the response to our get request
	ErrReadingResponse
)

type Error struct {
	Type ErrorType
	msg  string
}

func (se Error) Error() string {
	return se.msg
}

func NewError(t ErrorType, msg string) Error {
	return Error{Type: t, msg: msg}
}
