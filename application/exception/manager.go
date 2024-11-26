package applicationexception

import (
	"fmt"
	"reflect"
)

type Code string

const (
	InvalidArgumentCode  Code = "INVALID_ARGUMENT"  // Represents an invalid argument error.
	NotFoundCode         Code = "NOT_FOUND"         // Represents a not found error.
	AlreadyExistsCode    Code = "ALREADY_EXISTS"    // Represents an already exists error.
	PermissionDeniedCode Code = "PERMISSION_DENIED" // Represents a permission denied error.
	UnauthenticatedCode  Code = "UNAUTHENTICATED"   // Represents an unauthenticated error.
	InternalErrorCode    Code = "INTERNAL"          // Represents an internal error.
)

type Exception struct {
	Code    Code
	Message any
	Error   error
}

func (e *Exception) GetError() *string {
	if e.Error != nil {
		err := e.Error.Error()
		return &err
	}
	return nil
}

func (e *Exception) GetHttpCode() int {
	switch e.Code {
	case InvalidArgumentCode:
		return 400
	case NotFoundCode:
		return 404
	case AlreadyExistsCode:
		return 409
	case PermissionDeniedCode:
		return 403
	case UnauthenticatedCode:
		return 401
	case InternalErrorCode:
		return 500
	default:
		return 500
	}
}

func (e *Exception) IsEqual(err *Exception) bool {
	if err == nil {
		return e == nil
	}
	if e == nil {
		return err == nil
	}
	return e.Code == err.Code && reflect.DeepEqual(e.Message, err.Message)
}

func InvalidArgument(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    InvalidArgumentCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func NotFound(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    NotFoundCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func AlreadyExists(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    AlreadyExistsCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func PermissionDenied(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    PermissionDeniedCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func Unauthenticated(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    UnauthenticatedCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func Internal(message any, err error) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    InternalErrorCode,
		Message: errorMessage,
		Error:   err,
	}
}

func Conflict(message any) *Exception {
	errorMessage := convertToString(message)
	return &Exception{
		Code:    AlreadyExistsCode,
		Message: errorMessage,
		Error:   fmt.Errorf(errorMessage),
	}
}

func convertToString(message any) string {
	if msg, ok := message.(string); ok {
		return msg
	}
	return fmt.Sprintf("%v", message)
}
