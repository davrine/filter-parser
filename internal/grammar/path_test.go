package grammar

import (
	"testing"

	"github.com/di-wu/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	parser, err := ast.New([]byte("members[value eq \"2819c223-7f76-453a-919d-413861904646\"].displayName"))
	assert.NoError(t, err)
	node, err := Path(parser)
	assert.NoError(t, err)
	assert.Equal(t, "[\"Path\",[[\"ValuePath\",[[\"AttrPath\",[[\"AttrName\",\"members\"]]],[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"value\"]]],[\"CompareOp\",\"eq\"],[\"String\",\"\\\"2819c223-7f76-453a-919d-413861904646\\\"\"]]]]],[\"AttrName\",\"displayName\"]]]", node.String())
}
