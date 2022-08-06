package validator

import (
	"fmt"
	"testing"
)

func TestFieldError_Error_ReturnCorrectResult(t *testing.T) {
	params := []FieldError{
		{
			Name:   "arithmetic",
			Detail: "division by zero",
		},
		{
			Name:   "runtime",
			Detail: "info",
		},
	}
	for _, param := range params {
		if param.Error() != fmt.Sprintf("field - %s; detail: {%s}", param.Name, param.Detail) {
			t.Error("invalid error message")
		}
	}
}

func TestFieldErrors_Error_ReturnCorrectResult(t *testing.T) {
	sliceFieldErrors := []Errors{
		[]error{
			FieldError{
				Name:   "arithmetic",
				Detail: "division by zero",
			},
			FieldError{
				Name:   "runtime",
				Detail: "info",
			},
		},
		[]error{
			FieldError{
				Name:   "logic",
				Detail: "detail",
			},
			FieldError{
				Name:   "static",
				Detail: "line 123",
			},
		},
	}
	for _, fieldErrors := range sliceFieldErrors {
		var message string
		for index, fieldError := range fieldErrors {
			message += fieldError.Error()
			if index < len(fieldErrors)-1 {
				message += "\n"
			}
		}
		if fieldErrors.Error() != message {
			t.Errorf("invalid errors message")
		}
	}
}
