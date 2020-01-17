package helpers
import "errors"

var (
	// ErrParamMissing required field missing error
	 ErrInvalidRequest = errors.New("Invalid Request")
	 ErrParamMissing = errors.New("paramMissing")
	 ErrInvalidField = errors.New("Invalid Field")


)
