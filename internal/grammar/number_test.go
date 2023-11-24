package grammar

import (
	"testing"

	"github.com/di-wu/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {
	parser, err := ast.New([]byte("-10.0e+01"))
	assert.NoError(t, err)
	node, err := Number(parser)
	assert.NoError(t, err)
	assert.Equal(t, "[\"Number\",[[\"Minus\",\"-\"],[\"Int\",\"10\"],[\"Frac\",[[\"Digits\",\"0\"]]],[\"Exp\",[[\"Sign\",\"+\"],[\"Digits\",\"01\"]]]]]", node.String())
}
