package server

type ErrorHandle struct{}

var Err = &ErrorHandle{}

func (e *ErrorHandle) CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
