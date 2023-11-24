package grammar

import (
	"testing"

	"github.com/di-wu/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestFalse(t *testing.T) {
	example := "FaLSe"
	parser, err := ast.New([]byte(example))
	assert.NoError(t, err)
	node, err := False(parser)
	assert.NoError(t, err)
	assert.Equal(t, "False", node.TypeString())
	assert.Equal(t, example, node.Value)
}

func TestTrue(t *testing.T) {
	example := "TRue"
	parser, err := ast.New([]byte("TRue"))
	assert.NoError(t, err)
	node, err := True(parser)
	assert.NoError(t, err)
	assert.Equal(t, "True", node.TypeString())
	assert.Equal(t, example, node.Value)
}
