package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFile_WithInlineComments(t *testing.T) {
	// Inline comments are NOT stripped — value includes the comment text.
	// This matches standard .env behaviour where inline comments are ambiguous.
	path := writeTemp(t, "KEY=value # inline comment\n")
	env, err := ParseFile(path)
	require.NoError(t, err)
	assert.Equal(t, "value # inline comment", env["KEY"])
}

func TestParseFile_WhitespaceAroundEquals(t *testing.T) {
	path := writeTemp(t, "  KEY  =  value  \n")
	env, err := ParseFile(path)
	require.NoError(t, err)
	assert.Equal(t, "value", env["KEY"])
}

func TestParseFile_EmptyValue(t *testing.T) {
	path := writeTemp(t, "EMPTY=\n")
	env, err := ParseFile(path)
	require.NoError(t, err)
	assert.Equal(t, "", env["EMPTY"])
}

func TestParseFile_MultipleKeys(t *testing.T) {
	content := "A=1\nB=2\nC=3\n"
	path := writeTemp(t, content)
	env, err := ParseFile(path)
	require.NoError(t, err)
	assert.Len(t, env, 3)
	assert.Equal(t, "1", env["A"])
	assert.Equal(t, "2", env["B"])
	assert.Equal(t, "3", env["C"])
}

func TestParseFile_ValueWithEquals(t *testing.T) {
	path := writeTemp(t, "URL=http://example.com?foo=bar\n")
	env, err := ParseFile(path)
	require.NoError(t, err)
	assert.Equal(t, "http://example.com?foo=bar", env["URL"])
}
