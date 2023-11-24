package grammar

import (
	"testing"

	"github.com/di-wu/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestURI(t *testing.T) {
	example := "urn:ietf:params:scim:schemas:core:2.0:User:userName"
	parser, err := ast.New([]byte(example))
	assert.NoError(t, err)
	node, err := URI(parser)
	assert.NoError(t, err)
	assert.Equal(t, "URI", node.TypeString())
	assert.Equal(t, example, node.Value+"userName")
}
