package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAttrPath(t *testing.T) {
	attrPath := "urn:ietf:params:scim:schemas:core:2.0:User:name.familyName"
	parsed, err := ParseAttrPath([]byte(attrPath))
	assert.NoError(t, err)
	assert.Equal(t, attrPath, parsed.String())
}

func TestParseAttrPathErrors(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		_, err := ParseAttrPath([]byte(""))
		assert.Error(t, err)
	})

	t.Run("URN with no attribute", func(t *testing.T) {
		_, err := ParseAttrExp([]byte("urn:ietf:params:scim:schemas:core:2.0:"))
		assert.Error(t, err)
	})

	t.Run("URN with bad attribute", func(t *testing.T) {
		_, err := ParseAttrExp([]byte("urn:ietf:params:scim:schemas:core:2.0:Use"))
		assert.Error(t, err)
	})
}
