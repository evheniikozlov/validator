package validator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	tagName                        = "validator"
	tagRegexpExpression            = "(?P<name>\\w+)[ ]*:[ ]*(?P<argument>\\\"[^\\\".]*\\\"|\\w*)"
	validationNameRegexpExpression = "\\w+"
)

var tagRegexp, validationNameRegexp *regexp.Regexp

type validator struct {
	validations map[string]Validation
	templates   map[string]map[string]any
}

type validationTagEntity struct {
	name     string
	argument string
}

type validationWithArgument struct {
	validation Validation
	argument   string
}

func init() {
	tagRegexp = regexp.MustCompile(tagRegexpExpression)
	validationNameRegexp = regexp.MustCompile(validationNameRegexpExpression)
}

func NewValidatorWithPresets() validator {
	vld := NewValidator()
	vld.validations["min"] = ValidationMin{}
	vld.validations["max"] = ValidationMax{}
	vld.validations["min_length"] = ValidationMinLength{}
	vld.validations["max_length"] = ValidationMaxLength{}
	vld.validations["assert_false"] = ValidationAssertFalse{}
	vld.validations["assert_true"] = ValidationAssertTrue{}
	vld.validations["is_nil"] = ValidationNil{}
	vld.validations["not_nil"] = ValidationNotNil{}
	vld.validations["is_empty"] = ValidationEmpty{}
	vld.validations["not_empty"] = ValidationNotEmpty{}
	vld.validations["is_blank"] = ValidationBlank{}
	vld.validations["not_blank"] = ValidationNotBlank{}
	vld.validations["regexp_match"] = ValidationRegexpMatch{}
	return vld
}

func NewValidator() validator {
	var vld validator
	vld.validations = make(map[string]Validation)
	vld.templates = make(map[string]map[string]any)
	return vld
}

func (vld validator) SetValidation(name string, validation Validation) {
	if validationNameRegexp.MatchString(name) {
		vld.validations[name] = validation
	}
}

func (vld validator) DeleteValidation(name string) {
	delete(vld.validations, name)
}

func (vld validator) Compile(data any) error {
	dataValue := reflect.ValueOf(data)
	dataType := dataValue.Type()
	if dataType.Kind() != reflect.Struct {
		return errors.New("data type is not struct")
	}
	template, err := vld.unsafeTemplate(dataValue)
	if err != nil {
		return err
	}
	vld.templates[vld.templateName(dataType)] = template
	return nil
}

func (vld validator) MustCompile(data any) {
	if err := vld.Compile(data); err != nil {
		panic(err)
	}
}

func (vld validator) Validate(data any) error {
	dataValue := reflect.ValueOf(data)
	if dataValue.Type().Kind() != reflect.Struct {
		return errors.New("data type is not struct")
	}
	return vld.unsafeValidate(dataValue)
}

func (vld validator) templateName(tp reflect.Type) string {
	return fmt.Sprintf("%s/%s", tp.PkgPath(), tp.Name())
}

func (vld validator) unsafeTemplate(dataValue reflect.Value) (map[string]any, error) {
	var errs Errors
	template := make(map[string]any)
	for fieldIndex := 0; fieldIndex < dataValue.NumField(); fieldIndex++ {
		fieldValue := dataValue.Field(fieldIndex)
		fieldType := dataValue.Type().Field(fieldIndex)
		if fieldType.Type.Kind() == reflect.Struct {
			subTemplate, err := vld.unsafeTemplate(fieldValue)
			if err != nil {
				errs = append(errs, FieldError{Name: fieldType.Name, Detail: err.Error()})
			} else {
				template[fieldType.Name] = subTemplate
			}
			continue
		}
		validationTag, exist := fieldType.Tag.Lookup(tagName)
		if !exist {
			continue
		}
		validationTagEntities := vld.parseValidationTag(validationTag)
		for _, tagEntity := range validationTagEntities {
			validation, exist := vld.validations[tagEntity.name]
			if !exist {
				continue
			}
			err := validation.ValidateMeta(fieldValue.Interface(), tagEntity.argument)
			if err != nil {
				errs = append(errs, FieldError{Name: fieldType.Name, Detail: err.Error()})
				continue
			}
			if validations, ok := template[fieldType.Name].([]validationWithArgument); ok {
				template[fieldType.Name] = append(validations, validationWithArgument{validation: validation, argument: tagEntity.argument})
			} else {
				template[fieldType.Name] = []validationWithArgument{{validation: validation, argument: tagEntity.argument}}
			}
		}
	}
	if len(errs) != 0 {
		return nil, errs
	}
	return template, nil
}

func (vld validator) unsafeValidate(dataValue reflect.Value) error {
	templateName := vld.templateName(dataValue.Type())
	if template, exist := vld.templates[templateName]; exist {
		return vld.unsafeValidateWithTemplate(dataValue, template)
	} else {
		return errors.New(fmt.Sprintf("template %s is not exist", templateName))
	}
}

func (vld validator) unsafeValidateWithTemplate(dataValue reflect.Value, template map[string]any) error {
	var errs Errors
	for fieldName, value := range template {
		fieldValue := dataValue.FieldByName(fieldName)
		if validations, ok := value.([]validationWithArgument); ok {
			for _, validationWithArg := range validations {
				err := validationWithArg.validation.Validate(fieldValue.Interface(), validationWithArg.argument)
				if err != nil {
					errs = append(errs, FieldError{Name: fieldName, Detail: err.Error()})
				}
			}
		} else if subTemplate, ok := value.(map[string]any); ok {
			err := vld.unsafeValidateWithTemplate(fieldValue, subTemplate)
			if err != nil {
				errs = append(errs, FieldError{Name: fieldName, Detail: err.Error()})
			}
		}
	}
	if len(errs) != 0 {
		return errs
	}
	return nil
}

func (vld validator) parseValidationTag(tag string) []validationTagEntity {
	var validationTagEntities []validationTagEntity
	allGroups := tagRegexp.FindAllStringSubmatch(tag, -1)
	for _, group := range allGroups {
		validationTagEntities = append(validationTagEntities, validationTagEntity{
			name:     group[1],
			argument: strings.Trim(group[2], "\""),
		})
	}
	return validationTagEntities
}
