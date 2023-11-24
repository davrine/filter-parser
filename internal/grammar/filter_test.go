package grammar

import (
	"testing"

	"github.com/di-wu/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestFilterAnd(t *testing.T) {
	nodeStrings := []string{
		"[\"FilterOr\",[[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"title\"]]]]],[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"userType\"]]],[\"CompareOp\",\"eq\"],[\"String\",\"\\\"Employee\\\"\"]]]]]]]",
		"[\"FilterOr\",[[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"userType\"]]],[\"CompareOp\",\"eq\"],[\"String\",\"\\\"Employee\\\"\"]]],[\"ValuePath\",[[\"AttrPath\",[[\"AttrName\",\"emails\"]]],[\"ValueLogExpAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"type\"]]],[\"CompareOp\",\"eq\"],[\"String\",\"\\\"work\\\"\"]]],[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"value\"]]],[\"CompareOp\",\"co\"],[\"String\",\"\\\"@example.com\\\"\"]]]]]]]]]]]",
	}
	for i, example := range []string{
		"title pr and userType eq \"Employee\"",
		"userType eq \"Employee\" and emails[type eq \"work\" and value co \"@example.com\"]",
	} {
		t.Run(example, func(t *testing.T) {
			parser, err := ast.New([]byte(example))
			assert.NoError(t, err)
			node, err := Filter(parser)
			assert.NoError(t, err)
			assert.Equal(t, nodeStrings[i], node.String())
		})
	}
}

func TestFilterNot(t *testing.T) {
	example := "userType ne \"Employee\" and not (emails co \"example.com\" or emails.value co \"example.org\")"
	parser, err := ast.New([]byte(example))
	assert.NoError(t, err)
	node, err := Filter(parser)
	assert.NoError(t, err)
	assert.Equal(t, "[\"FilterOr\",[[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"userType\"]]],[\"CompareOp\",\"ne\"],[\"String\",\"\\\"Employee\\\"\"]]],[\"FilterNot\",[[\"FilterPrecedence\",[[\"FilterOr\",[[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"emails\"]]],[\"CompareOp\",\"co\"],[\"String\",\"\\\"example.com\\\"\"]]]]],[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"emails\"],[\"AttrName\",\"value\"]]],[\"CompareOp\",\"co\"],[\"String\",\"\\\"example.org\\\"\"]]]]]]]]]]]]]]]", node.String())
}

func TestFilterOr(t *testing.T) {
	nodeStrings := []string{
		"[\"FilterOr\",[[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"title\"]]]]]]],[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"userType\"]]],[\"CompareOp\",\"eq\"],[\"String\",\"\\\"Intern\\\"\"]]]]]]]",
		"[\"FilterOr\",[[\"FilterAnd\",[[\"ValuePath\",[[\"AttrPath\",[[\"AttrName\",\"emails\"]]],[\"ValueLogExpAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"type\"]]],[\"CompareOp\",\"eq\"],[\"String\",\"\\\"work\\\"\"]]],[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"value\"]]],[\"CompareOp\",\"co\"],[\"String\",\"\\\"@example.com\\\"\"]]]]]]]]],[\"FilterAnd\",[[\"ValuePath\",[[\"AttrPath\",[[\"AttrName\",\"ims\"]]],[\"ValueLogExpAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"type\"]]],[\"CompareOp\",\"eq\"],[\"String\",\"\\\"xmpp\\\"\"]]],[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"value\"]]],[\"CompareOp\",\"co\"],[\"String\",\"\\\"@foo.com\\\"\"]]]]]]]]]]]",
	}
	for i, example := range []string{
		"title pr or userType eq \"Intern\"",
		"emails[type eq \"work\" and value co \"@example.com\"] or ims[type eq \"xmpp\" and value co \"@foo.com\"]",
	} {
		t.Run(example, func(t *testing.T) {
			parser, err := ast.New([]byte(example))
			assert.NoError(t, err)
			node, err := Filter(parser)
			assert.NoError(t, err)
			assert.Equal(t, nodeStrings[i], node.String())
		})
	}
}

func TestFilterParentheses(t *testing.T) {
	nodeStrings := []string{
		"[\"FilterOr\",[[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"userType\"]]],[\"CompareOp\",\"eq\"],[\"String\",\"\\\"Employee\\\"\"]]],[\"FilterPrecedence\",[[\"FilterOr\",[[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"emails\"],[\"AttrName\",\"type\"]]],[\"CompareOp\",\"eq\"],[\"String\",\"\\\"work\\\"\"]]]]]]]]]]]]]",
		"[\"FilterOr\",[[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"userType\"]]],[\"CompareOp\",\"eq\"],[\"String\",\"\\\"Employee\\\"\"]]],[\"FilterPrecedence\",[[\"FilterOr\",[[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"emails\"]]],[\"CompareOp\",\"co\"],[\"String\",\"\\\"example.com\\\"\"]]]]],[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"emails\"],[\"AttrName\",\"value\"]]],[\"CompareOp\",\"co\"],[\"String\",\"\\\"example.org\\\"\"]]]]]]]]]]]]]",
		"[\"FilterOr\",[[\"FilterAnd\",[[\"FilterNot\",[[\"FilterPrecedence\",[[\"FilterOr\",[[\"FilterAnd\",[[\"AttrExp\",[[\"AttrPath\",[[\"AttrName\",\"emails\"]]],[\"CompareOp\",\"co\"],[\"String\",\"\\\"example.com\\\"\"]]]]]]]]]]]]]]]",
	}
	for i, example := range []string{
		"userType eq \"Employee\" and (emails.type eq \"work\")",
		"userType eq \"Employee\" and (emails co \"example.com\" or emails.value co \"example.org\")",
		"not (emails co \"example.com\")",
	} {
		t.Run(example, func(t *testing.T) {
			parser, err := ast.New([]byte(example))
			assert.NoError(t, err)
			node, err := Filter(parser)
			assert.NoError(t, err)
			assert.Equal(t, nodeStrings[i], node.String())
		})
	}
}
