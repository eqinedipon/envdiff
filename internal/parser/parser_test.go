package parser

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func writeTemp(t *testing.T, content string) string {
	t.Helper()
	tmp := filepath.Join(t.TempDir(), ".env")
	err := os.WriteFile(tmp, []byte(content), 0644)
	require.NoError(t, err)
	return tmp
}

func TestParseFile_Basic(t *testing.T) {
	path := writeTemp(t, "APP_ENV=production\nDB_HOST=localhost\n")
	env, err := ParseFile(path)
	require.NoError(t, err)
	assert.Equal(t, EnvMap{"APP_ENV": "production", "DB_HOST": "localhost"}, env)
}

func TestParseFile_SkipsCommentsAndBlanks(t *testing.T) {
	content := "# comment\n\nKEY=value\n"
	path := writeTemp(t, content)
	env, err := ParseFile(path)
	require.NoError(t, err)
	assert.Equal(t, EnvMap{"KEY": "value"}, env)
}

func TestParseFile_StripQuotes(t *testing.T) {
	path := writeTemp(t, `SINGLE='hello'\nDOUBLE="world"\n`)
	_ = path
	// Test stripQuotes directly
	assert.Equal(t, "hello", stripQuotes("'hello'"))
	assert.Equal(t, "world", stripQuotes(`"world"`))
	assert.Equal(t, "noquotes", stripQuotes("noquotes"))
}

func TestParseFile_MalformedLine(t *testing.T) {
	path := writeTemp(t, "BADLINE\n")
	_, err := ParseFile(path)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "malformed line")
}

func TestParseFile_FileNotFound(t *testing.T) {
	_, err := ParseFile("/nonexistent/.env")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "opening file")
}

func TestParseFile_EmptyKey(t *testing.T) {
	path := writeTemp(t, "=value\n")
	_, err := ParseFile(path)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "empty key")
}
