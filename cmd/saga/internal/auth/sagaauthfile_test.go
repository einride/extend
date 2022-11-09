package auth

import (
	"encoding/json"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestSagaAuthFile_MarshalJSONEmpty(t *testing.T) {
	expected := strings.TrimSpace(`
{
  "clientID": "",
  "clientSecret": ""
}
	`)
	var authFile SagaAuthFile
	assert.NilError(t, json.Unmarshal([]byte(expected), &authFile))
	actual, err := json.MarshalIndent(&authFile, "", "  ")
	assert.NilError(t, err)
	assert.Equal(t, expected, string(actual))
}

func TestSagaAuthFile_MarshalJSON(t *testing.T) {
	expected := strings.TrimSpace(`
{
  "clientID": "abc",
  "clientSecret": "def"
}
	`)
	var authFile SagaAuthFile
	assert.NilError(t, json.Unmarshal([]byte(expected), &authFile))
	actual, err := json.MarshalIndent(&authFile, "", "  ")
	assert.NilError(t, err)
	assert.Equal(t, expected, string(actual))
}
