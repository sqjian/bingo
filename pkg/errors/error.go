package errors

import "fmt"

const (
	UnknownCode = -1
)

var (
	TestErr              = New(1, "Unknown")
	PluginMethodNotFound = New(2, "NewPlugin not found in plugin")
	ServerMethodNotFound = New(3, "NewServer not found in acceptor")
)

type Error interface {
	Error() string
	ErrorCode() int
	ErrorMsg() string
}

type ErrorPredefined struct {
	Msg  string
	Code int
}

func New(code int, msg string) *ErrorPredefined {
	return &ErrorPredefined{
		Code: code,
		Msg:  msg,
	}
}

func NewWithErr(err error) Error {
	if err == nil {
		panic("internal error")
	}
	return New(UnknownCode, err.Error())
}
func (e *ErrorPredefined) Error() string {
	return fmt.Sprintf("error: %s | Code: %d", e.Msg, e.Code)
}
func (e *ErrorPredefined) ErrorCode() int {
	return e.Code
}
func (e *ErrorPredefined) ErrorMsg() string {
	return e.Msg
}
