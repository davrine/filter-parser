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
