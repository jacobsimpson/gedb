package lexer

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectStar(t *testing.T) {
	assert := assert.New(t)
	l := NewLexer(strings.NewReader("SELECT * FROM users"))
	assert.NotNil(l)

	token, err := l.Token()
	assert.NoError(err)
	assert.NotNil(token)
	assert.Equal(SelectToken, token.Type)
	assert.Equal("SELECT", token.Raw)
}
