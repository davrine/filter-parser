package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePathAttrPath(t *testing.T) {
	for _, attrPath := range []string{
		"members",
		"name.familyName",
	} {
		t.Run(attrPath, func(t *testing.T) {
			parsed, err := ParsePath([]byte(attrPath))
			assert.NoError(t, err)
			assert.Equal(t, attrPath, parsed.String())
		})
	}
}
func TestParsePathValuePath(t *testing.T) {
	for _, valuePath := range []string{
		"members[value eq \"2819c223-7f76-453a-919d-413861904646\"]",
		"members[value eq \"2819c223-7f76-453a-919d-413861904646\"].displayName",
	} {
		t.Run(valuePath, func(t *testing.T) {
			parsed, err := ParsePath([]byte(valuePath))
			assert.NoError(t, err)
			assert.Equal(t, valuePath, parsed.String())
		})
	}
}

func TestParsePath(t *testing.T) {
	for _, path := range []string{
		"members",
		"name.familyName",
		"addresses[type eq \"work\"]",
		"members[value eq \"2819c223-7f76-453a-919d-413861904646\"]",
		"members[value eq \"2819c223-7f76-453a-919d-413861904646\"].displayName",
	} {
		t.Run(path, func(t *testing.T) {
			parsed, err := ParsePath([]byte(path))
			assert.NoError(t, err)
			assert.Equal(t, path, parsed.String())
		})
	}
}

func TestParsePathErrors(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		_, err := ParsePath([]byte(""))
		assert.Error(t, err)
	})

	t.Run("Invalid path", func(t *testing.T) {
		_, err := ParsePath([]byte("members a"))
		assert.Error(t, err)
	})
}
