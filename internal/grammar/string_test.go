package grammar

import (
	"testing"

	"github.com/di-wu/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	example := "\"2819c223-7f76-453a-919d-413861904646\""
	parser, err := ast.New([]byte(example))
	assert.NoError(t, err)
	node, err := String(parser)
	assert.NoError(t, err)
	assert.Equal(t, "String", node.TypeString())
	assert.Equal(t, example, node.Value)
}

func TestStringComplex(t *testing.T) {
	example := "\"W/\\\"990-6468886345120203448\\\"\""
	parser, err := ast.New([]byte(example))
	assert.NoError(t, err)
	node, err := String(parser)
	assert.NoError(t, err)
	assert.Equal(t, "String", node.TypeString())
	assert.Equal(t, example, node.Value)
}
