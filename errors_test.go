package filter

import (
	"strconv"
	"testing"

	typ "github.com/scim2/filter-parser/v2/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestInvalidChildTypeError(t *testing.T) {
	var invalidChildType *internalError
	parentTyp, invalidType := 1, 2
	err := invalidChildTypeError(parentTyp, invalidType)
	assert.ErrorAs(t, err, &invalidChildType)
	expected := "internal error: invalid child type for " + typ.Stringer[parentTyp] + " (00" + strconv.Itoa(parentTyp) + "): " + typ.Stringer[invalidType] + " (00" + strconv.Itoa(invalidType) + ")"
	assert.Equal(t, expected, err.Error())
}

func TestInvalidLengthError(t *testing.T) {
	var invalidLength *internalError
	parentTyp, len, actual := 1, 2, 3
	err := invalidLengthError(parentTyp, len, actual)
	assert.ErrorAs(t, err, &invalidLength)
	expected := "internal error: length was not equal to " + strconv.Itoa(len) + " for " + typ.Stringer[parentTyp] + " (00" + strconv.Itoa(parentTyp) + "), got " + strconv.Itoa(actual) + " elements"
	assert.Equal(t, expected, err.Error())
}

func TestInvalidTypeError(t *testing.T) {
	var invalidType *internalError
	parentTyp, actual := 1, 2
	err := invalidTypeError(parentTyp, actual)
	assert.ErrorAs(t, err, &invalidType)
	expected := "internal error: invalid type: expected " + typ.Stringer[parentTyp] + " (00" + strconv.Itoa(parentTyp) + "), actual " + typ.Stringer[actual] + " (00" + strconv.Itoa(actual) + ")"
	assert.Equal(t, expected, err.Error())
}

func TestMissingValueError(t *testing.T) {
	var missingValue *internalError
	parentTyp, actual := 1, 2
	err := missingValueError(parentTyp, actual)
	assert.ErrorAs(t, err, &missingValue)
	expected := "internal error: missing a required value for " + typ.Stringer[parentTyp] + " (00" + strconv.Itoa(parentTyp) + "): " + typ.Stringer[actual] + " (00" + strconv.Itoa(actual) + ")"
	assert.Equal(t, expected, err.Error())
}
