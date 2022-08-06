package validator

import "fmt"

type FieldError struct {
	Name   string
	Detail string
}

type Errors []error

func (err FieldError) Error() string {
	return fmt.Sprintf("field - %s; detail: {%s}", err.Name, err.Detail)
}

func (errs Errors) Error() string {
	var errsMessage string
	for index, err := range errs {
		fmt.Println(err)
		errsMessage += err.Error()
		if index < len(errs)-1 {
			errsMessage += "\n"
		}
	}
	return errsMessage
}
