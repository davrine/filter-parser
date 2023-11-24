package filter

import (
	"testing"

	"github.com/di-wu/parser/ast"
	"github.com/stretchr/testify/assert"
)

func TestParseFilterAnd(t *testing.T) {
	filter := "title pr and userType eq \"Employee\""
	expression, err := ParseFilter([]byte(filter))
	assert.NoError(t, err)
	exported, err := Export(expression)
	assert.NoError(t, err)
	assert.Equal(t, filter, exported)
}

func TestParseFilterAttrExp(t *testing.T) {
	filter := "schemas eq \"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User\""
	expression, err := ParseFilter([]byte(filter))
	assert.NoError(t, err)
	exported, err := Export(expression)
	assert.NoError(t, err)
	assert.Equal(t, filter, exported)
}

func TestParseFilterCaseInsensitivity(t *testing.T) {
	filter := "NAME pr and not (FIRST eq \"test\") and ANOTHER ne \"test\""
	expression, err := ParseFilter([]byte(filter))
	assert.NoError(t, err)
	exported, err := Export(expression)
	assert.NoError(t, err)
	assert.Equal(t, filter, exported)
}

func TestParseFilterNot(t *testing.T) {
	filter := "not (emails co \"example.com\" or emails.value co \"example.org\")"
	expression, err := ParseFilter([]byte(filter))
	assert.NoError(t, err)
	exported, err := Export(expression)
	assert.NoError(t, err)
	assert.Equal(t, filter, exported)
}

func TestParseFilterOr(t *testing.T) {
	filter := "title pr or userType eq \"Intern\""
	expression, err := ParseFilter([]byte(filter))
	assert.NoError(t, err)
	exported, err := Export(expression)
	assert.NoError(t, err)
	assert.Equal(t, filter, exported)
}

func TestParseFilterParentheses(t *testing.T) {
	filter := "(emails.type eq \"work\")"
	expression, err := ParseFilter([]byte(filter))
	assert.NoError(t, err)
	exported, err := Export(expression)
	assert.NoError(t, err)
	assert.Equal(t, filter, exported)
}

func TestParseFilterValuePath(t *testing.T) {
	filter := "emails[type eq \"work\" and value co \"@example.com\"]"
	expression, err := ParseFilter([]byte(filter))
	assert.NoError(t, err)
	exported, err := Export(expression)
	assert.NoError(t, err)
	assert.Equal(t, filter, exported)
}

func TestParseFilter(t *testing.T) {
	for _, example := range []string{
		"userName eq \"bjensen\"",
		"name.familyName co \"O'Malley\"",
		"userName sw \"J\"",
		"urn:ietf:params:scim:schemas:core:2.0:User:userName sw \"J\"",
		"title pr",
		"meta.lastModified gt \"2011-05-13T04:42:34Z\"",
		"meta.lastModified ge \"2011-05-13T04:42:34Z\"",
		"meta.lastModified lt \"2011-05-13T04:42:34Z\"",
		"meta.lastModified le \"2011-05-13T04:42:34Z\"",
		"title pr and userType eq \"Employee\"",
		"title pr or userType eq \"Intern\"",
		"schemas eq \"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User\"",
		"userType eq \"Employee\" and (emails co \"example.com\" or emails.value co \"example.org\")",
		"userType ne \"Employee\" and not (emails co \"example.com\" or emails.value co \"example.org\")",
		"userType eq \"Employee\" and (emails.type eq \"work\")",
		"userType eq \"Employee\" and emails[type eq \"work\" and value co \"@example.com\"]",
		"emails[type eq \"work\" and value co \"@example.com\"] or ims[type eq \"xmpp\" and value co \"@foo.com\"]",

		"name pr and userName pr and title pr",
		"name pr and not (first eq \"test\") and another ne \"test\"",
		"NAME pr and not (FIRST eq \"test\") and ANOTHER ne \"test\"",
		"name pr or userName pr or title pr",
	} {
		t.Run(example, func(t *testing.T) {
			ast, err := ParseFilter([]byte(example))
			assert.NoError(t, err)
			exported, err := Export(ast)
			assert.NoError(t, err)
			assert.Equal(t, example, exported)
		})
	}
}

func TestParseFilterErrors(t *testing.T) {
	var newNode ast.Node
	var p config
	var internalError *internalError

	t.Run("Empty string", func(t *testing.T) {
		_, err := ParseFilter([]byte(""))
		assert.Error(t, err)
	})

	t.Run("Invalid filter without pr", func(t *testing.T) {
		_, err := ParseFilter([]byte("name eq"))
		assert.Error(t, err)
	})

	t.Run("Invalid filter with pr", func(t *testing.T) {
		_, err := ParseFilter([]byte("name pr \"a\""))
		assert.Error(t, err)
	})

	t.Run("filterAnd wrong type", func(t *testing.T) {
		newNode.Type = 1
		_, err := p.parseFilterAnd(&newNode)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &internalError)
	})

	t.Run("filterAnd no children", func(t *testing.T) {
		newNode.Type = 2
		_, err := p.parseFilterAnd(&newNode)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &internalError)
	})

	t.Run("filterOr wrong type", func(t *testing.T) {
		newNode.Type = 2
		_, err := p.parseFilterOr(&newNode)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &internalError)
	})

	t.Run("filterOr no children", func(t *testing.T) {
		newNode.Type = 1
		_, err := p.parseFilterOr(&newNode)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &internalError)
	})

	t.Run("filterValue ValuePath", func(t *testing.T) {
		newNode.Type = 10
		_, err := p.parseFilterValue(&newNode)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &internalError)
	})

	t.Run("filterValue AttrExp", func(t *testing.T) {
		newNode.Type = 6
		_, err := p.parseFilterValue(&newNode)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &internalError)
	})

	t.Run("filterValue FilterNot", func(t *testing.T) {
		newNode.Type = 3
		_, err := p.parseFilterValue(&newNode)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &internalError)
	})

	t.Run("filterValue FilterPrecedence", func(t *testing.T) {
		newNode.Type = 4
		_, err := p.parseFilterValue(&newNode)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &internalError)
	})

	t.Run("filterValue FilterOr", func(t *testing.T) {
		newNode.Type = 1
		_, err := p.parseFilterValue(&newNode)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &internalError)
	})

	t.Run("filterValue Default", func(t *testing.T) {
		newNode.Type = 2
		_, err := p.parseFilterValue(&newNode)
		assert.Error(t, err)
		assert.ErrorAs(t, err, &internalError)
	})
}
