package mediators

import (
	"fmt"
	"strings"
)

// ##################### Individual error #####################

// ErrorType defines the kind of erros for the MediatorError
type ErrorType string

const (
	Authentication ErrorType = "authentication"
	Conflict       ErrorType = "conflict"
	Missing        ErrorType = "missing"
	Undefined      ErrorType = "undefined"
	Validation     ErrorType = "validation"
	NotFound       ErrorType = "notFound"
	DbError        ErrorType = "db_error"
)

type MediatorError struct {
	Code        string
	Description string
	Type        ErrorType
}

var errorList = []MediatorError{
	{"err_not_found", "Not found", Missing},
	{"err_unprocessable_entity", "Unprocessable entity", Conflict},
	{"err_internal_server_error", "%s", Undefined},
	{"err_invalid_data", "Invalid Data", Validation},
	{"err_unauthorized", "unauthorized", Authentication},
	{"not_found", "not found", NotFound},
}

// implement `Error() string` from the built-in error interface
func (e *MediatorError) Error() string {
	return e.Description
}

// GetType returns the error type
func (e *MediatorError) GetType() ErrorType {
	return e.Type
}

// GetError returns an error of type ErrorMediator
func (e *MediatorError) GetError(code string, value interface{}) *MediatorError {
	if err, ok := contains(errorList, code); ok {
		return &MediatorError{err.Code, setDesc(err.Description, value), err.Type}
	}
	return &MediatorError{"internal_server_error", "Internal Server Error", Undefined}
}

func NewMediatorError() *MediatorError {
	return &MediatorError{"no_error", "", Undefined}
}

// ##################### Helper functions #####################

func setDesc(desc string, value interface{}) string {
	switch value.(type) {
	case string:
		if strings.Contains(desc, "%") {
			return fmt.Sprintf(desc, value)
		}
		return desc
	default:
		return desc
	}
}

func contains(list []MediatorError, value string) (MediatorError, bool) {
	var emptyError MediatorError
	for _, i := range list {
		if i.Code == value {
			return i, true
		}
	}
	return emptyError, false
}
