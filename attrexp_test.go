package filter

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/di-wu/parser/ast"
	"github.com/scim2/filter-parser/v2/internal/grammar"
	"github.com/stretchr/testify/assert"
)

func TestParseAttrExpPr(t *testing.T) {
	attrExp := "userName pr"
	parsed, err := ParseAttrExp([]byte(attrExp))
	assert.NoError(t, err)
	assert.Equal(t, attrExp, parsed.String())
}

func TestParseAttrExpSw(t *testing.T) {
	attrExp := "userName sw \"J\""
	parsed, err := ParseAttrExp([]byte(attrExp))
	assert.NoError(t, err)
	assert.Equal(t, attrExp, parsed.String())
}

func TestParseNumber(t *testing.T) {
	for _, test := range []struct {
		nStr     string
		expected interface{}
	}{
		{
			nStr:     "-5.1e-2",
			expected: -0.051,
		},
		{
			nStr:     "-5.1e2",
			expected: float64(-510),
		},
		{
			nStr:     "-510",
			expected: -510,
		},
	} {
		t.Run(test.nStr, func(t *testing.T) {
			p, _ := ast.New([]byte(test.nStr))
			n, err := grammar.Number(p)
			if err != nil {
				t.Fatal(err)
			}
			{ // Empty config.
				i, err := config{}.parseNumber(n)
				if err != nil {
					t.Fatal(err)
				}
				if i != test.expected {
					t.Error(test.expected, i)
				}
			}
			{ // Config with useNumber = true.
				d := json.NewDecoder(strings.NewReader(test.nStr))
				d.UseNumber()
				var number json.Number
				if err := d.Decode(&number); err != nil {
					t.Error(err)
				}

				i, err := config{
					useNumber: true,
				}.parseNumber(n)
				if err != nil {
					t.Fatal(err)
				}
				if i != json.Number(test.nStr) {
					t.Error(test.nStr, i)
				}

				// Check if equal to json.Decode.
				if i != number {
					t.Error(number, i)
				}
			}
			{ // Config with useNumber = true.
				d := json.NewDecoder(strings.NewReader(test.nStr))
				d.UseNumber()
				var number json.Number
				if err := d.Decode(&number); err != nil {
					t.Error(err)
				}

				i, err := config{
					useNumber: true,
				}.parseNumber(n)
				if err != nil {
					t.Fatal(err)
				}
				if i != json.Number(test.nStr) {
					t.Error(test.nStr, i)
				}

				// Check if equal to json.Decode.
				if i != number {
					t.Error(number, i)
				}
			}
		})
	}
}

func TestParseAttrExpErrors(t *testing.T) {

	t.Run("Empty string", func(t *testing.T) {
		_, err := ParseAttrExp([]byte(""))
		assert.Error(t, err)
	})

	t.Run("Invalid attrExp not pr operator", func(t *testing.T) {
		_, err := ParseAttrExp([]byte("userName eq "))
		assert.Error(t, err)
	})

	t.Run("Invalid attrExp pr operator", func(t *testing.T) {
		_, err := ParseAttrExp([]byte("userName pr \"a\""))
		assert.Error(t, err)
	})
}
