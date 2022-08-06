package validator

import "testing"

func TestValidationMin_Validate_ReturnNil_ValidParams(t *testing.T) {
	params := []struct {
		value int
		min   string
	}{
		{
			value: 2,
			min:   "1",
		},
		{
			value: 20,
			min:   "10",
		},
	}
	for _, param := range params {
		if err := (ValidationMin{}.Validate(param.value, param.min)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationMin_Validate_ReturnError_InvalidParams(t *testing.T) {
	params := []struct {
		value any
		min   string
	}{
		{
			value: 1,
			min:   "2",
		},
		{
			value: 10,
			min:   "20",
		},
		{
			value: "str",
			min:   "20",
		},
		{
			value: 10,
			min:   "str",
		},
	}
	for _, param := range params {
		if err := (ValidationMin{}.Validate(param.value, param.min)); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationMin_IsCorrect_ReturnNil_ValidParams(t *testing.T) {
	params := []struct {
		value int
		min   string
	}{
		{
			value: 1,
			min:   "2",
		},
		{
			value: 20,
			min:   "10",
		},
	}
	for _, param := range params {
		if err := (ValidationMin{}.ValidateMeta(param.value, param.min)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationMin_IsCorrect_ReturnError_InvalidParams(t *testing.T) {
	params := []struct {
		value any
		min   string
	}{
		{
			value: "str",
			min:   "20",
		},
		{
			value: 10,
			min:   "str",
		},
	}
	for _, param := range params {
		if err := (ValidationMin{}.ValidateMeta(param.value, param.min)); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationMax_Validate_ReturnNil_ValidParams(t *testing.T) {
	params := []struct {
		value int
		max   string
	}{
		{
			value: 1,
			max:   "2",
		},
		{
			value: 10,
			max:   "20",
		},
	}
	for _, param := range params {
		if err := (ValidationMax{}.Validate(param.value, param.max)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationMax_Validate_ReturnError_InvalidParams(t *testing.T) {
	params := []struct {
		value any
		max   string
	}{
		{
			value: 2,
			max:   "1",
		},
		{
			value: 20,
			max:   "10",
		},
		{
			value: "str",
			max:   "20",
		},
		{
			value: 10,
			max:   "str",
		},
	}
	for _, param := range params {
		if err := (ValidationMax{}.Validate(param.value, param.max)); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationMax_IsCorrect_ReturnNil_ValidParams(t *testing.T) {
	params := []struct {
		value int
		max   string
	}{
		{
			value: 2,
			max:   "1",
		},
		{
			value: 10,
			max:   "20",
		},
	}
	for _, param := range params {
		if err := (ValidationMax{}.ValidateMeta(param.value, param.max)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationMax_IsCorrect_ReturnError_InvalidParams(t *testing.T) {
	params := []struct {
		value any
		max   string
	}{
		{
			value: "str",
			max:   "20",
		},
		{
			value: 10,
			max:   "str",
		},
	}
	for _, param := range params {
		if err := (ValidationMax{}.ValidateMeta(param.value, param.max)); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationMinLength_Validate_ReturnNil_ValidParams(t *testing.T) {
	params := []struct {
		value string
		min   string
	}{
		{
			value: "str",
			min:   "3",
		},
		{
			value: "st",
			min:   "2",
		},
	}
	for _, param := range params {
		if err := (ValidationMinLength{}.Validate(param.value, param.min)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationMinLength_Validate_ReturnError_InvalidParams(t *testing.T) {
	params := []struct {
		value any
		min   string
	}{
		{
			value: "str",
			min:   "4",
		},
		{
			value: "st",
			min:   "3",
		},
		{
			value: 1,
			min:   "20",
		},
		{
			value: "str",
			min:   "str",
		},
	}
	for _, param := range params {
		if err := (ValidationMinLength{}.Validate(param.value, param.min)); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationMinLength_IsCorrect_ReturnNil_ValidParams(t *testing.T) {
	params := []struct {
		value string
		min   string
	}{
		{
			value: "str",
			min:   "5",
		},
		{
			value: "st",
			min:   "2",
		},
	}
	for _, param := range params {
		if err := (ValidationMinLength{}.ValidateMeta(param.value, param.min)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationMinLength_IsCorrect_ReturnError_InvalidParams(t *testing.T) {
	params := []struct {
		value any
		min   string
	}{
		{
			value: 1,
			min:   "20",
		},
		{
			value: "str",
			min:   "str",
		},
	}
	for _, param := range params {
		if err := (ValidationMinLength{}.ValidateMeta(param.value, param.min)); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationMaxLength_Validate_ReturnNil_ValidParams(t *testing.T) {
	params := []struct {
		value string
		max   string
	}{
		{
			value: "str",
			max:   "5",
		},
		{
			value: "st",
			max:   "2",
		},
	}
	for _, param := range params {
		if err := (ValidationMaxLength{}.Validate(param.value, param.max)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationMaxLength_Validate_ReturnError_InvalidParams(t *testing.T) {
	params := []struct {
		value any
		max   string
	}{
		{
			value: "str",
			max:   "2",
		},
		{
			value: "st",
			max:   "1",
		},
		{
			value: 1,
			max:   "20",
		},
		{
			value: "str",
			max:   "str",
		},
	}
	for _, param := range params {
		if err := (ValidationMaxLength{}.Validate(param.value, param.max)); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationMaxLength_IsCorrect_ReturnNil_ValidParams(t *testing.T) {
	params := []struct {
		value string
		max   string
	}{
		{
			value: "str",
			max:   "2",
		},
		{
			value: "st",
			max:   "5",
		},
	}
	for _, param := range params {
		if err := (ValidationMaxLength{}.ValidateMeta(param.value, param.max)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationMaxLength_IsCorrect_ReturnError_InvalidParams(t *testing.T) {
	params := []struct {
		value any
		max   string
	}{
		{
			value: 1,
			max:   "20",
		},
		{
			value: "str",
			max:   "str",
		},
	}
	for _, param := range params {
		if err := (ValidationMaxLength{}.ValidateMeta(param.value, param.max)); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationAssertFalse_Validate_ReturnNil_ValidValue(t *testing.T) {
	if err := (ValidationAssertFalse{}.Validate(false, "")); err != nil {
		t.Error(err)
	}
}

func TestValidationAssertFalse_Validate_ReturnError_InvalidValue(t *testing.T) {
	values := []any{true, -1, -10, "str"}
	for _, value := range values {
		if err := (ValidationAssertFalse{}.Validate(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationAssertFalse_IsCorrect_ReturnNil_ValidValue(t *testing.T) {
	values := []bool{false, true}
	for _, value := range values {
		if err := (ValidationAssertFalse{}.ValidateMeta(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationAssertFalse_IsCorrect_ReturnError_InvalidValue(t *testing.T) {
	values := []any{-1, -10, "str"}
	for _, value := range values {
		if err := (ValidationAssertFalse{}.ValidateMeta(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationAssertTrue_Validate_ReturnNil_ValidValue(t *testing.T) {
	if err := (ValidationAssertTrue{}.Validate(true, "")); err != nil {
		t.Error(err)
	}
}

func TestValidationAssertTrue_Validate_ReturnError_InvalidValue(t *testing.T) {
	values := []any{false, -1, -10, "str"}
	for _, value := range values {
		if err := (ValidationAssertTrue{}.Validate(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationAssertTrue_IsCorrect_ReturnNil_ValidValue(t *testing.T) {
	values := []bool{false, true}
	for _, value := range values {
		if err := (ValidationAssertTrue{}.ValidateMeta(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationAssertTrue_IsCorrect_ReturnError_InvalidValue(t *testing.T) {
	values := []any{-1, -10, "str"}
	for _, value := range values {
		if err := (ValidationAssertTrue{}.ValidateMeta(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationNil_Validate_ReturnNil_ValidValue(t *testing.T) {
	if err := (ValidationNil{}.Validate(nil, "")); err != nil {
		t.Error(err)
	}
}

func TestValidationNil_Validate_ReturnError_InvalidValue(t *testing.T) {
	values := []any{-1, -10, "str"}
	for _, value := range values {
		if err := (ValidationNil{}.Validate(value, "")); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationNil_IsCorrect_ReturnNil_Always(t *testing.T) {
	values := []any{-1, -10, "str"}
	for _, value := range values {
		if err := (ValidationNil{}.ValidateMeta(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationNotNil_Validate_ReturnNil_ValidValue(t *testing.T) {
	values := []any{-1, -10, "str"}
	for _, value := range values {
		if err := (ValidationNotNil{}.Validate(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationNotNil_Validate_ReturnError_InvalidValue(t *testing.T) {
	if err := (ValidationNotNil{}.Validate(nil, "")); err == nil {
		t.Error("error is nil")
	}
}

func TestValidationNotNil_IsCorrect_ReturnNil_Always(t *testing.T) {
	values := []any{-1, -10, "str"}
	for _, value := range values {
		if err := (ValidationNotNil{}.ValidateMeta(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationEmpty_Validate_ReturnNil_ValidValue(t *testing.T) {
	if err := (ValidationEmpty{}.Validate("", "")); err != nil {
		t.Error(err)
	}
}

func TestValidationEmpty_Validate_ReturnError_InvalidValue(t *testing.T) {
	values := []any{-1, -10, "str"}
	for _, value := range values {
		if err := (ValidationEmpty{}.Validate(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationEmpty_IsCorrect_ReturnNil_ValidValue(t *testing.T) {
	values := []string{"", "str"}
	for _, value := range values {
		if err := (ValidationEmpty{}.ValidateMeta(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationEmpty_IsCorrect_ReturnError_InvalidValue(t *testing.T) {
	values := []any{-1, -10}
	for _, value := range values {
		if err := (ValidationEmpty{}.ValidateMeta(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationNotEmpty_Validate_ReturnNil_ValidValue(t *testing.T) {
	values := []string{"   ", "str"}
	for _, value := range values {
		if err := (ValidationNotEmpty{}.Validate(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationNotEmpty_Validate_ReturnError_InvalidValue(t *testing.T) {
	values := []any{"", 1, 10}
	for _, value := range values {
		if err := (ValidationNotEmpty{}.Validate(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationNotEmpty_IsCorrect_ReturnNil_ValidValue(t *testing.T) {
	values := []string{"", "str"}
	for _, value := range values {
		if err := (ValidationNotEmpty{}.ValidateMeta(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationNotEmpty_IsCorrect_ReturnError_InvalidValue(t *testing.T) {
	values := []any{-1, -10}
	for _, value := range values {
		if err := (ValidationNotEmpty{}.ValidateMeta(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationBlank_Validate_ReturnNil_ValidValue(t *testing.T) {
	values := []string{"", "   "}
	for _, value := range values {
		if err := (ValidationBlank{}.Validate(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationBlank_Validate_ReturnError_InvalidValue(t *testing.T) {
	values := []any{-1, -10, "str"}
	for _, value := range values {
		if err := (ValidationBlank{}.Validate(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationBlank_IsCorrect_ReturnNil_ValidValue(t *testing.T) {
	values := []string{"   ", "str"}
	for _, value := range values {
		if err := (ValidationBlank{}.ValidateMeta(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationBlank_IsCorrect_ReturnError_InvalidValue(t *testing.T) {
	values := []any{-1, -10}
	for _, value := range values {
		if err := (ValidationBlank{}.ValidateMeta(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationNotBlank_Validate_ReturnNil_ValidValue(t *testing.T) {
	values := []string{"st", "str"}
	for _, value := range values {
		if err := (ValidationNotBlank{}.Validate(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationNotBlank_Validate_ReturnError_InvalidValue(t *testing.T) {
	values := []any{-1, -10, "  "}
	for _, value := range values {
		if err := (ValidationNotBlank{}.Validate(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationNotBlank_IsCorrect_ReturnNil_ValidValue(t *testing.T) {
	values := []string{"   ", "str"}
	for _, value := range values {
		if err := (ValidationNotBlank{}.ValidateMeta(value, "")); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationNotBlank_IsCorrect_ReturnError_InvalidValue(t *testing.T) {
	values := []any{-1, -10}
	for _, value := range values {
		if err := (ValidationNotBlank{}.ValidateMeta(value, "")); err == nil {
			t.Error("error is nil")
		}
	}
}

func TestValidationRegexpMatch_Validate_ReturnNil_ValidParams(t *testing.T) {
	params := []struct {
		value  string
		regexp string
	}{
		{
			value:  "123",
			regexp: "\\d+",
		},
		{
			value:  "abc",
			regexp: "\\w+",
		},
	}
	for _, param := range params {
		if err := (ValidationRegexpMatch{}.Validate(param.value, param.regexp)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationRegexpMatch_Validate_ReturnError_InvalidParams(t *testing.T) {
	params := []struct {
		value  any
		regexp string
	}{
		{
			value:  "---",
			regexp: "\\w+",
		},
		{
			value:  "abc",
			regexp: "\\d+",
		},
		{
			value:  1,
			regexp: "20",
		},
		{
			value:  "abc",
			regexp: "a(b",
		},
	}
	for _, param := range params {
		if err := (ValidationRegexpMatch{}.Validate(param.value, param.regexp)); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationRegexpMatch_IsCorrect_ReturnNil_ValidParams(t *testing.T) {
	params := []struct {
		value  string
		regexp string
	}{
		{
			value:  "123",
			regexp: "\\w+",
		},
		{
			value:  "abc",
			regexp: "\\d+",
		},
	}
	for _, param := range params {
		if err := (ValidationRegexpMatch{}.ValidateMeta(param.value, param.regexp)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationRegexpMatch_IsCorrect_ReturnError_InvalidParams(t *testing.T) {
	params := []struct {
		value  any
		regexp string
	}{
		{
			value:  "---",
			regexp: "a(b",
		},
		{
			value:  1,
			regexp: "20",
		},
	}
	for _, param := range params {
		if err := (ValidationRegexpMatch{}.ValidateMeta(param.value, param.regexp)); err == nil {
			t.Error("err is nil")
		}
	}
}

func TestValidationMin_Validate_ReturnNil_NumberTypes(t *testing.T) {
	params := []struct {
		value any
		min   string
	}{
		{
			value: float64(2),
			min:   "1",
		},
		{
			value: float32(2),
			min:   "1",
		},
		{
			value: int64(2),
			min:   "1",
		},
		{
			value: int32(2),
			min:   "1",
		},
		{
			value: 2,
			min:   "1",
		},
		{
			value: uint64(2),
			min:   "1",
		},
		{
			value: uint32(2),
			min:   "1",
		},
		{
			value: uint(2),
			min:   "1",
		},
	}
	for _, param := range params {
		if err := (ValidationMin{}.Validate(param.value, param.min)); err != nil {
			t.Error(err)
		}
	}
}

func TestValidationMin_Validate_ReturnError_InvalidTypes(t *testing.T) {
	params := []struct {
		value any
		min   string
	}{
		{
			false,
			"2",
		},
		{
			true,
			"2",
		},
	}
	for _, param := range params {
		if err := (ValidationMin{}.Validate(param.value, param.min)); err == nil {
			t.Error("error is nil")
		}
	}
}
