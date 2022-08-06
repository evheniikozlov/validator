package validator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Validation interface {
	ValidateMeta(value any, argument string) error
	Validate(value any, argument string) error
}

type ValidationMin struct {
}

func (_ ValidationMin) ValidateMeta(fieldValue any, validationArgument string) error {
	var errs Errors
	_, err := getFloat(fieldValue)
	if err != nil {
		errs = append(errs, err)
	}
	_, err = strconv.ParseFloat(validationArgument, 64)
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func (_ ValidationMin) Validate(fieldValue any, validationArgument string) error {
	floatFieldValue, err := getFloat(fieldValue)
	if err != nil {
		return err
	}
	minValue, err := strconv.ParseFloat(validationArgument, 64)
	if err != nil {
		return err
	}
	if floatFieldValue < minValue {
		return errors.New("value less than min")
	}
	return nil
}

type ValidationMax struct {
}

func (_ ValidationMax) ValidateMeta(fieldValue any, validationArgument string) error {
	var errs Errors
	_, err := getFloat(fieldValue)
	if err != nil {
		errs = append(errs, err)
	}
	_, err = strconv.ParseFloat(validationArgument, 64)
	if err != nil {
		errs = append(errs, err)
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func (_ ValidationMax) Validate(fieldValue any, validationArgument string) error {
	floatFieldValue, err := getFloat(fieldValue)
	if err != nil {
		return err
	}
	maxValue, err := strconv.ParseFloat(validationArgument, 64)
	if err != nil {
		return err
	}
	if floatFieldValue > maxValue {
		return errors.New("value more than max")
	}
	return nil
}

type ValidationMinLength struct {
}

func (_ ValidationMinLength) ValidateMeta(fieldValue any, validationArgument string) error {
	var errs Errors
	_, err := strconv.ParseInt(validationArgument, 10, 64)
	if err != nil {
		errs = append(errs, err)
	}
	if _, ok := fieldValue.(string); !ok {
		errs = append(errs, errors.New("value type is not string"))
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func (_ ValidationMinLength) Validate(fieldValue any, validationArgument string) error {
	minValueLength, err := strconv.ParseInt(validationArgument, 10, 64)
	if err != nil {
		return err
	}
	if stringFieldValue, ok := fieldValue.(string); ok {
		if int64(len(stringFieldValue)) < minValueLength {
			return errors.New("value length less than min")
		}
	} else {
		return errors.New("value type is not string")
	}
	return nil
}

type ValidationMaxLength struct {
}

func (_ ValidationMaxLength) ValidateMeta(fieldValue any, validationArgument string) error {
	var errs Errors
	_, err := strconv.ParseInt(validationArgument, 10, 64)
	if err != nil {
		errs = append(errs, err)
	}
	if _, ok := fieldValue.(string); !ok {
		errs = append(errs, errors.New("value type is not string"))
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func (_ ValidationMaxLength) Validate(fieldValue any, validationArgument string) error {
	maxValueLength, err := strconv.ParseInt(validationArgument, 10, 64)
	if err != nil {
		return err
	}
	if stringFieldValue, ok := fieldValue.(string); ok {
		if int64(len(stringFieldValue)) > maxValueLength {
			return errors.New("value length more than max")
		}
	} else {
		return errors.New("value type is not string")
	}
	return nil
}

type ValidationAssertFalse struct {
}

func (_ ValidationAssertFalse) ValidateMeta(fieldValue any, validationArgument string) error {
	if _, ok := fieldValue.(bool); !ok {
		return errors.New("value type is not bool")
	}
	return nil
}

func (_ ValidationAssertFalse) Validate(fieldValue any, validationArgument string) error {
	if boolFieldValue, ok := fieldValue.(bool); ok {
		if boolFieldValue {
			return errors.New("value is true")
		}
	} else {
		return errors.New("value type is not bool")
	}
	return nil
}

type ValidationAssertTrue struct {
}

func (_ ValidationAssertTrue) ValidateMeta(fieldValue any, validationArgument string) error {
	if _, ok := fieldValue.(bool); !ok {
		return errors.New("value type is not bool")
	}
	return nil
}

func (_ ValidationAssertTrue) Validate(fieldValue any, validationArgument string) error {
	if boolFieldValue, ok := fieldValue.(bool); ok {
		if !boolFieldValue {
			return errors.New("value is false")
		}
	} else {
		return errors.New("value type is not bool")
	}
	return nil
}

type ValidationNil struct {
}

func (_ ValidationNil) ValidateMeta(fieldValue any, validationArgument string) error {
	return nil
}

func (_ ValidationNil) Validate(fieldValue any, validationArgument string) error {
	if fieldValue != nil {
		return errors.New("value is nil")
	}
	return nil
}

type ValidationNotNil struct {
}

func (_ ValidationNotNil) ValidateMeta(fieldValue any, validationArgument string) error {
	return nil
}

func (_ ValidationNotNil) Validate(fieldValue any, validationArgument string) error {
	if fieldValue == nil {
		return errors.New("value is nil")
	}
	return nil
}

type ValidationEmpty struct {
}

func (_ ValidationEmpty) ValidateMeta(fieldValue any, validationArgument string) error {
	if _, ok := fieldValue.(string); !ok {
		return errors.New("value type is not string")
	}
	return nil
}

func (_ ValidationEmpty) Validate(fieldValue any, validationArgument string) error {
	if stringFieldValue, ok := fieldValue.(string); ok {
		if stringFieldValue != "" {
			return errors.New("value is not empty")
		}
	} else {
		return errors.New("value type is not string")
	}
	return nil
}

type ValidationNotEmpty struct {
}

func (_ ValidationNotEmpty) ValidateMeta(fieldValue any, validationArgument string) error {
	if _, ok := fieldValue.(string); !ok {
		return errors.New("value type is not string")
	}
	return nil
}

func (_ ValidationNotEmpty) Validate(fieldValue any, validationArgument string) error {
	if stringFieldValue, ok := fieldValue.(string); ok {
		if stringFieldValue == "" {
			return errors.New("value is empty")
		}
	} else {
		return errors.New("value type is not string")
	}
	return nil
}

type ValidationBlank struct {
}

func (_ ValidationBlank) ValidateMeta(fieldValue any, validationArgument string) error {
	if _, ok := fieldValue.(string); !ok {
		return errors.New("value type is not string")
	}
	return nil
}

func (_ ValidationBlank) Validate(fieldValue any, validationArgument string) error {
	if stringFieldValue, ok := fieldValue.(string); ok {
		if strings.TrimSpace(stringFieldValue) != "" {
			return errors.New("value is not ValidationBlank")
		}
	} else {
		return errors.New("value type is not string")
	}
	return nil
}

type ValidationNotBlank struct {
}

func (_ ValidationNotBlank) ValidateMeta(fieldValue any, validationArgument string) error {
	if _, ok := fieldValue.(string); !ok {
		return errors.New("value type is not string")
	}
	return nil
}

func (_ ValidationNotBlank) Validate(fieldValue any, validationArgument string) error {
	if stringFieldValue, ok := fieldValue.(string); ok {
		if strings.TrimSpace(stringFieldValue) == "" {
			return errors.New("value is ValidationBlank")
		}
	} else {
		return errors.New("value type is not string")
	}
	return nil
}

type ValidationRegexpMatch struct {
}

func (_ ValidationRegexpMatch) ValidateMeta(fieldValue any, validationArgument string) error {
	if stringFieldValue, ok := fieldValue.(string); ok {
		if _, err := regexp.MatchString(validationArgument, stringFieldValue); err != nil {
			return err
		}
	} else {
		return errors.New("value type is not string")
	}
	return nil
}

func (_ ValidationRegexpMatch) Validate(fieldValue any, validationArgument string) error {
	if stringFieldValue, ok := fieldValue.(string); ok {
		if match, err := regexp.MatchString(validationArgument, stringFieldValue); err == nil {
			if !match {
				return errors.New("value is not match regexp")
			}
		} else {
			return err
		}
	} else {
		return errors.New("value type is not string")
	}
	return nil
}

func getFloat(data any) (float64, error) {
	switch typedData := data.(type) {
	case float64:
		return typedData, nil
	case float32:
		return float64(typedData), nil
	case int64:
		return float64(typedData), nil
	case int32:
		return float64(typedData), nil
	case int:
		return float64(typedData), nil
	case uint64:
		return float64(typedData), nil
	case uint32:
		return float64(typedData), nil
	case uint:
		return float64(typedData), nil
	case string:
		return strconv.ParseFloat(typedData, 64)
	default:
		return 0, errors.New("invalid field type")
	}
}
