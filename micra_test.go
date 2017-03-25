package micra

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateLengthIs10(t *testing.T) {
	s, _ := Generate(10)

	assert.Len(t, s, 10)
}

func TestItErrorsWhenLengthIsLessThanFive(t *testing.T) {
	_, err := Generate(4)

	if assert.Error(t, err) {
		assert.Equal(t, "Length must be more or equal to 5", err.Error())
	}
}

func TestItErrorsWhenLengthIsLargerThanDefaultGlyphs(t *testing.T) {
	_, err := Generate(63)

	if assert.Error(t, err) {
		assert.Equal(t, "Length of ID cannot be longer than glyphs (62)", err.Error())
	}
}

func TestAddingCustomGlyphs(t *testing.T) {
	s, _ := Generate(5, "aaaaa")

	if assert.Len(t, s, 5) {
		assert.Equal(t, "aaaaa", s)
	}
}

func TestGenerateInt(t *testing.T) {
	s, err := GenerateInt(5)

	if assert.NoError(t, err) {
		assert.IsType(t, new(int64), &s)
	}
}

func TestGenerateIntErrors(t *testing.T) {
	_, err := GenerateInt(4)

	assert.Error(t, err)
}
