package helpers

import "errors"

var (
	// ErrParamMissing required field missing error
	ErrInvalidRequest = errors.New("invalid request")
	ErrParamMissing   = errors.New("paramMissing")
	ErrInvalidField   = errors.New("invalid field")
	ErrSomeThingWentWrng   = errors.New("something went wrong")
)
