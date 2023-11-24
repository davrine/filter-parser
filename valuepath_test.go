package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValuePath(t *testing.T) {
	for _, valuePath := range []string{
		"emails[type eq \"work\"]",
		"emails[not (type eq \"work\")]",
		"emails[type eq \"work\" and value co \"@example.com\"]",
	} {
		parsed, err := ParseValuePath([]byte(valuePath))
		assert.NoError(t, err)
		assert.Equal(t, valuePath, parsed.String())
	}
}

func TestParseValuePathErrors(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		_, err := ParseValuePath([]byte(""))
		assert.Error(t, err)
	})

	t.Run("Empty values", func(t *testing.T) {
		_, err := ParseValuePath([]byte("emails[]"))
		assert.Error(t, err)
	})

	t.Run("Empty bad AttrExp", func(t *testing.T) {
		_, err := ParseValuePath([]byte("emails[type eq]"))
		assert.Error(t, err)
	})
}
