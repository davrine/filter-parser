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
