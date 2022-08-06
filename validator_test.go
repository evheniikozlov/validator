package validator

import (
	"reflect"
	"testing"
)

const validationsCount = 13

type validTestingStructure struct {
	Age     int    `validator:"min: 18; max: 80"`
	Name    string `validator:"min_length: 4; max_length: 16; not_exist: not_exist"`
	Content string
	Info    struct {
		Money       int `validator:"min: 1000"`
		HomeAddress string
	}
}

type invalidTestingStructure struct {
	Age  string `validator:"min: 18; max: 80"`
	Name int    `validator:"min_length: 4; max_length: 16"`
	Info struct {
		Money string `validator:"min: 1000"`
	}
}

func Test_NewValidator_ReturnValidatorWithZeroValidations(t *testing.T) {
	vld := NewValidator()
	if validationsLength := len(vld.validations); validationsLength != 0 {
		t.Errorf("validations length is %d but expected 0", validationsLength)
	}
}

func Test_NewValidatorWithPresets_ReturnValidatorWithAllValidations(t *testing.T) {
	vld := NewValidatorWithPresets()
	if validationsLength := len(vld.validations); validationsLength != validationsCount {
		t.Errorf("validations length is %d but expected %d", validationsLength, validationsCount)
	}
}

func TestValidator_SetValidation_SetValidation_ValidName(t *testing.T) {
	const validationName = "min"
	vld := NewValidator()
	vld.SetValidation(validationName, ValidationMin{})
	if _, exist := vld.validations[validationName]; !exist {
		t.Error("not set validation with valid name")
	}
}

func TestValidator_SetValidation_NotSetValidation_InvalidName(t *testing.T) {
	const validationName = "---"
	vld := NewValidator()
	vld.SetValidation(validationName, ValidationMin{})
	if _, exist := vld.validations[validationName]; exist {
		t.Error("set validation with invalid name")
	}
}

func TestValidator_DeleteValidation_DeleteValidation_All(t *testing.T) {
	const validationName = "min"
	vld := NewValidator()
	vld.SetValidation(validationName, ValidationMin{})
	vld.DeleteValidation(validationName)
	if _, exist := vld.validations[validationName]; exist {
		t.Error("deleted validation exist")
	}
}

func TestValidator_Compile_ReturnNilAndSetTemplate_ValidStructure(t *testing.T) {
	vld := NewValidatorWithPresets()
	err := vld.Compile(validTestingStructure{})
	if err != nil {
		t.Error(err)
	}
	if _, exist := vld.templates[vld.templateName(reflect.TypeOf(validTestingStructure{}))]; !exist {
		t.Error("template is not exist")
	}
}

func TestValidator_Compile_ReturnErrorAndNotSetTemplate_InvalidStructure(t *testing.T) {
	vld := NewValidatorWithPresets()
	err := vld.Compile(invalidTestingStructure{})
	if err == nil {
		t.Error("error is nil")
	}
	if _, exist := vld.templates[vld.templateName(reflect.TypeOf(validTestingStructure{}))]; exist {
		t.Error("template is exist")
	}
}

func TestValidator_Compile_ReturnError_NotStructure(t *testing.T) {
	vld := NewValidator()
	err := vld.Compile(123)
	if err == nil {
		t.Error("error is nil")
	}
}

func TestValidator_MustCompile_NotPanic_ValidStructure(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error(err)
		}
	}()
	vld := NewValidatorWithPresets()
	vld.MustCompile(validTestingStructure{})
}

func TestValidator_MustCompile_Panic_InvalidStructure(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("no panic")
		}
	}()
	vld := NewValidatorWithPresets()
	vld.MustCompile(invalidTestingStructure{})
}

func TestValidator_MustCompile_Panic_NotStructure(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("no panic")
		}
	}()
	vld := NewValidator()
	vld.MustCompile(123)
}

func TestValidator_Validate_ReturnNil_ValidData(t *testing.T) {
	data := validTestingStructure{
		Age:  20,
		Name: "Susan",
		Info: struct {
			Money       int `validator:"min: 1000"`
			HomeAddress string
		}{
			Money: 3000,
		},
	}
	vld := NewValidatorWithPresets()
	vld.MustCompile(data)
	err := vld.Validate(data)
	if err != nil {
		t.Error(err)
	}
}

func TestValidator_Validate_ReturnError_InvalidData(t *testing.T) {
	data := validTestingStructure{
		Age:  10,
		Name: "bob",
		Info: struct {
			Money       int `validator:"min: 1000"`
			HomeAddress string
		}{
			Money: 300,
		},
	}
	vld := NewValidatorWithPresets()
	vld.MustCompile(data)
	err := vld.Validate(data)
	if err == nil {
		t.Error("error is nil")
	}
}

func TestValidator_Validate_ReturnError_NotStructure(t *testing.T) {
	vld := NewValidatorWithPresets()
	err := vld.Validate(123)
	if err == nil {
		t.Error("error is nil")
	}
}

func TestValidator_Validate_ReturnError_TemplateNotSet(t *testing.T) {
	data := validTestingStructure{
		Age:  20,
		Name: "Susan",
		Info: struct {
			Money       int `validator:"min: 1000"`
			HomeAddress string
		}{
			Money: 3000,
		},
	}
	vld := NewValidatorWithPresets()
	err := vld.Validate(data)
	if err == nil {
		t.Error("error is nil")
	}
}
